package port

import (
	"context"

	agendav1 "go-challenge-agenda/gen/agenda/v1"
)

// AgendaPort is the API's port to the agenda service. Usecases depend on this interface
// instead of the concrete gRPC client, enabling tests and alternative implementations.
type AgendaPort interface {
	GetAvailability(ctx context.Context, req *agendav1.GetAvailabilityRequest) (*agendav1.GetAvailabilityResponse, error)
	CreateReservation(ctx context.Context, req *agendav1.CreateReservationRequest) (*agendav1.CreateReservationResponse, error)
	CancelReservation(ctx context.Context, req *agendav1.CancelReservationRequest) (*agendav1.CancelReservationResponse, error)
	GetReservation(ctx context.Context, req *agendav1.GetReservationRequest) (*agendav1.GetReservationResponse, error)
	ListReservations(ctx context.Context, req *agendav1.ListReservationsRequest) (*agendav1.ListReservationsResponse, error)
	ListDoctors(ctx context.Context, req *agendav1.ListDoctorsRequest) (*agendav1.ListDoctorsResponse, error)
	GetDoctor(ctx context.Context, req *agendav1.GetDoctorRequest) (*agendav1.GetDoctorResponse, error)
	ListPatients(ctx context.Context, req *agendav1.ListPatientsRequest) (*agendav1.ListPatientsResponse, error)
	GetPatient(ctx context.Context, req *agendav1.GetPatientRequest) (*agendav1.GetPatientResponse, error)
	CreatePatient(ctx context.Context, req *agendav1.CreatePatientRequest) (*agendav1.CreatePatientResponse, error)
	UpdatePatient(ctx context.Context, req *agendav1.UpdatePatientRequest) (*agendav1.UpdatePatientResponse, error)
	DeletePatient(ctx context.Context, req *agendav1.DeletePatientRequest) (*agendav1.DeletePatientResponse, error)
}
