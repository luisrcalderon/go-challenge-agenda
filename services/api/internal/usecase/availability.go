package usecase

import (
	"context"
	"fmt"

	agendav1 "go-challenge-agenda/gen/agenda/v1"
	"go-challenge-agenda/services/api/internal/domain"
)

// AvailabilityUsecase is coupled to the concrete gRPC client type.
// Candidates should extract an AgendaPort interface and inject it.
type AvailabilityUsecase struct {
	agendaClient agendav1.AgendaServiceClient
}

func NewAvailabilityUsecase(client agendav1.AgendaServiceClient) *AvailabilityUsecase {
	return &AvailabilityUsecase{agendaClient: client}
}

func (u *AvailabilityUsecase) GetAvailability(ctx context.Context, doctorID, date, resType string) (*domain.AvailabilityResponse, error) {
	pbType := agendav1.ReservationType_RESERVATION_TYPE_FOLLOW_UP
	if resType == "first_visit" {
		pbType = agendav1.ReservationType_RESERVATION_TYPE_FIRST_VISIT
	}

	resp, err := u.agendaClient.GetAvailability(ctx, &agendav1.GetAvailabilityRequest{
		DoctorId:        doctorID,
		Date:            date,
		ReservationType: pbType,
	})
	if err != nil {
		return nil, fmt.Errorf("agenda.GetAvailability: %w", err)
	}

	result := &domain.AvailabilityResponse{}
	for _, s := range resp.Slots {
		result.Slots = append(result.Slots, domain.AvailableSlot{StartsAt: s.StartsAt, EndsAt: s.EndsAt})
	}
	for _, r := range resp.FreeRanges {
		result.FreeRanges = append(result.FreeRanges, domain.TimeRange{From: r.From, To: r.To})
	}
	return result, nil
}
