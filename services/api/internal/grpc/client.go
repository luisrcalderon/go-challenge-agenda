package grpc

import (
	"fmt"

	agendav1 "go-challenge-agenda/gen/agenda/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewAgendaClient creates a gRPC client connected to the agenda service.
// NOTE: this returns a concrete type — the api usecase depends on it directly.
// Candidates should decouple this behind an interface.
func NewAgendaClient(addr string) (agendav1.AgendaServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("dial agenda service: %w", err)
	}
	return agendav1.NewAgendaServiceClient(conn), conn, nil
}
