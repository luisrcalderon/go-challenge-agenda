package usecase

import (
	"context"
	"fmt"

	agendav1 "go-challenge-agenda/gen/agenda/v1"
	"go-challenge-agenda/services/api/internal/domain"
	"go-challenge-agenda/services/api/internal/port"
)

type UserUsecase struct {
	agenda port.AgendaPort
}

func NewUserUsecase(agenda port.AgendaPort) *UserUsecase {
	return &UserUsecase{agenda: agenda}
}

func (u *UserUsecase) List(ctx context.Context) ([]domain.UserResponse, error) {
	resp, err := u.agenda.ListPatients(ctx, &agendav1.ListPatientsRequest{})
	if err != nil {
		return nil, fmt.Errorf("agenda.ListPatients: %w", err)
	}
	users := make([]domain.UserResponse, len(resp.Patients))
	for i, p := range resp.Patients {
		users[i] = protoPatientToDTO(p)
	}
	return users, nil
}

func (u *UserUsecase) Get(ctx context.Context, id string) (*domain.UserResponse, error) {
	resp, err := u.agenda.GetPatient(ctx, &agendav1.GetPatientRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("agenda.GetPatient: %w", err)
	}
	dto := protoPatientToDTO(resp.Patient)
	return &dto, nil
}

func (u *UserUsecase) Create(ctx context.Context, req *domain.CreateUserRequest) (*domain.UserResponse, error) {
	resp, err := u.agenda.CreatePatient(ctx, &agendav1.CreatePatientRequest{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("agenda.CreatePatient: %w", err)
	}
	dto := protoPatientToDTO(resp.Patient)
	return &dto, nil
}

func (u *UserUsecase) Update(ctx context.Context, id string, req *domain.UpdateUserRequest) (*domain.UserResponse, error) {
	// First fetch the existing patient to merge fields
	existing, err := u.agenda.GetPatient(ctx, &agendav1.GetPatientRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("agenda.GetPatient: %w", err)
	}

	name := existing.Patient.Name
	phone := existing.Patient.Phone
	email := existing.Patient.Email
	if req.Name != "" {
		name = req.Name
	}
	if req.Phone != "" {
		phone = req.Phone
	}
	if req.Email != "" {
		email = req.Email
	}

	resp, err := u.agenda.UpdatePatient(ctx, &agendav1.UpdatePatientRequest{
		Id: id, Name: name, Phone: phone, Email: email,
	})
	if err != nil {
		return nil, fmt.Errorf("agenda.UpdatePatient: %w", err)
	}
	dto := protoPatientToDTO(resp.Patient)
	return &dto, nil
}

func (u *UserUsecase) Delete(ctx context.Context, id string) error {
	_, err := u.agenda.DeletePatient(ctx, &agendav1.DeletePatientRequest{Id: id})
	return err
}

// ListReservations returns all reservations for the given user (patient) id.
func (u *UserUsecase) ListReservations(ctx context.Context, userID string) ([]domain.ReservationResponse, error) {
	resp, err := u.agenda.ListReservations(ctx, &agendav1.ListReservationsRequest{
		PatientId: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("agenda.ListReservations: %w", err)
	}
	out := make([]domain.ReservationResponse, len(resp.Reservations))
	for i, r := range resp.Reservations {
		out[i] = *reservationProtoToDTO(r)
	}
	return out, nil
}

func reservationProtoToDTO(r *agendav1.Reservation) *domain.ReservationResponse {
	if r == nil {
		return nil
	}
	typeStr := "follow_up"
	if r.Type == agendav1.ReservationType_RESERVATION_TYPE_FIRST_VISIT {
		typeStr = "first_visit"
	}
	statusStr := "confirmed"
	if r.Status == agendav1.ReservationStatus_RESERVATION_STATUS_CANCELLED {
		statusStr = "cancelled"
	}
	return &domain.ReservationResponse{
		ID: r.Id, DoctorID: r.DoctorId, PatientID: r.PatientId,
		StartsAt: r.StartsAt, EndsAt: r.EndsAt, Type: typeStr, Status: statusStr,
	}
}

func protoPatientToDTO(p *agendav1.Patient) domain.UserResponse {
	return domain.UserResponse{ID: p.Id, Name: p.Name, Phone: p.Phone, Email: p.Email}
}
