package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	agendav1 "go-challenge-agenda/gen/agenda/v1"
	"go-challenge-agenda/services/agenda/config"
	"go-challenge-agenda/services/agenda/internal/domain"
	agendagrpc "go-challenge-agenda/services/agenda/internal/grpc"
	"go-challenge-agenda/services/agenda/internal/repository/postgres"
	"go-challenge-agenda/services/agenda/internal/repository/sqlite"
	"go-challenge-agenda/services/agenda/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}

	switch cfg.DBDriver {
	case "sqlite3":
		if err := sqlite.Migrate(db); err != nil {
			log.Fatalf("migrate: %v", err)
		}
		if err := sqlite.Seed(db); err != nil {
			log.Fatalf("seed: %v", err)
		}
	case "postgres":
		if err := postgres.Migrate(db); err != nil {
			log.Fatalf("migrate: %v", err)
		}
		if err := postgres.Seed(db); err != nil {
			log.Fatalf("seed: %v", err)
		}
	}

	var doctorRepo domain.DoctorRepository
	var patientRepo domain.PatientRepository
	var reservationRepo domain.ReservationRepository
	var blockedSlotRepo domain.BlockedSlotRepository

	switch cfg.DBDriver {
	case "sqlite3":
		doctorRepo = sqlite.NewDoctorRepository(db)
		patientRepo = sqlite.NewPatientRepository(db)
		reservationRepo = sqlite.NewReservationRepository(db)
		blockedSlotRepo = sqlite.NewBlockedSlotRepository(db)
	case "postgres":
		doctorRepo = postgres.NewDoctorRepository(db)
		patientRepo = postgres.NewPatientRepository(db)
		reservationRepo = postgres.NewReservationRepository(db)
		blockedSlotRepo = postgres.NewBlockedSlotRepository(db)
	}

	availUC := usecase.NewAvailabilityUsecase(doctorRepo, reservationRepo, blockedSlotRepo)
	reservationUC := usecase.NewReservationUsecase(reservationRepo, patientRepo)
	blockedSlotUC := usecase.NewBlockedSlotUsecase(blockedSlotRepo)

	srv := agendagrpc.NewServer(doctorRepo, availUC, reservationUC, blockedSlotUC, patientRepo)

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpcLoggingInterceptor))
	agendav1.RegisterAgendaServiceServer(grpcServer, srv)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("agenda gRPC server listening on %s", cfg.GRPCAddr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("gRPC serve stopped: %v", err)
		}
	}()

	<-quit
	log.Println("shutting down agenda service...")

	grpcServer.GracefulStop()

	if sqlDB, err := db.DB(); err == nil {
		sqlDB.Close()
	}

	log.Println("agenda service stopped")
}

func grpcLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(start)
	code := codes.OK
	if err != nil {
		if st, ok := status.FromError(err); ok {
			code = st.Code()
		} else {
			code = codes.Unknown
		}
	}
	slog.Info("grpc_request",
		"method", info.FullMethod,
		"code", code.String(),
		"duration_ms", duration.Milliseconds(),
	)
	return resp, err
}

func openDB(cfg config.Config) (*gorm.DB, error) {
	switch cfg.DBDriver {
	case "sqlite3":
		return sqlite.Open(cfg.DBSource)
	case "postgres":
		return postgres.Open(cfg.DBSource)
	default:
		return nil, fmt.Errorf("unknown DB driver: %s", cfg.DBDriver)
	}
}
