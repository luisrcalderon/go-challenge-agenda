package grpc

import (
	"context"
	"fmt"

	agendav1 "go-challenge-agenda/gen/agenda/v1"
	"go-challenge-agenda/services/api/internal/port"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewAgendaClient creates a gRPC client connected to the agenda service.
func NewAgendaClient(addr string) (agendav1.AgendaServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("dial agenda service: %w", err)
	}
	return agendav1.NewAgendaServiceClient(conn), conn, nil
}

// agendaPortAdapter adapts the gRPC client (which has ...grpc.CallOption in every method) to port.AgendaPort.
type agendaPortAdapter struct {
	client agendav1.AgendaServiceClient
}

// NewAgendaPort returns an AgendaPort that delegates to the given gRPC client.
func NewAgendaPort(client agendav1.AgendaServiceClient) port.AgendaPort {
	return &agendaPortAdapter{client: client}
}

func (a *agendaPortAdapter) GetAvailability(ctx context.Context, req *agendav1.GetAvailabilityRequest) (*agendav1.GetAvailabilityResponse, error) {
	return a.client.GetAvailability(ctx, req)
}
func (a *agendaPortAdapter) CreateReservation(ctx context.Context, req *agendav1.CreateReservationRequest) (*agendav1.CreateReservationResponse, error) {
	return a.client.CreateReservation(ctx, req)
}
func (a *agendaPortAdapter) CancelReservation(ctx context.Context, req *agendav1.CancelReservationRequest) (*agendav1.CancelReservationResponse, error) {
	return a.client.CancelReservation(ctx, req)
}
func (a *agendaPortAdapter) GetReservation(ctx context.Context, req *agendav1.GetReservationRequest) (*agendav1.GetReservationResponse, error) {
	return a.client.GetReservation(ctx, req)
}
func (a *agendaPortAdapter) ListReservations(ctx context.Context, req *agendav1.ListReservationsRequest) (*agendav1.ListReservationsResponse, error) {
	return a.client.ListReservations(ctx, req)
}
func (a *agendaPortAdapter) ListDoctors(ctx context.Context, req *agendav1.ListDoctorsRequest) (*agendav1.ListDoctorsResponse, error) {
	return a.client.ListDoctors(ctx, req)
}
func (a *agendaPortAdapter) GetDoctor(ctx context.Context, req *agendav1.GetDoctorRequest) (*agendav1.GetDoctorResponse, error) {
	return a.client.GetDoctor(ctx, req)
}
func (a *agendaPortAdapter) ListPatients(ctx context.Context, req *agendav1.ListPatientsRequest) (*agendav1.ListPatientsResponse, error) {
	return a.client.ListPatients(ctx, req)
}
func (a *agendaPortAdapter) GetPatient(ctx context.Context, req *agendav1.GetPatientRequest) (*agendav1.GetPatientResponse, error) {
	return a.client.GetPatient(ctx, req)
}
func (a *agendaPortAdapter) CreatePatient(ctx context.Context, req *agendav1.CreatePatientRequest) (*agendav1.CreatePatientResponse, error) {
	return a.client.CreatePatient(ctx, req)
}
func (a *agendaPortAdapter) UpdatePatient(ctx context.Context, req *agendav1.UpdatePatientRequest) (*agendav1.UpdatePatientResponse, error) {
	return a.client.UpdatePatient(ctx, req)
}
func (a *agendaPortAdapter) DeletePatient(ctx context.Context, req *agendav1.DeletePatientRequest) (*agendav1.DeletePatientResponse, error) {
	return a.client.DeletePatient(ctx, req)
}
