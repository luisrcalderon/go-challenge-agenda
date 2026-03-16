package usecase

import (
	"context"
	"fmt"

	agendav1 "go-challenge-agenda/gen/agenda/v1"
	"go-challenge-agenda/services/api/internal/domain"
	"go-challenge-agenda/services/api/internal/port"
)

type AvailabilityUsecase struct {
	agenda port.AgendaPort
}

func NewAvailabilityUsecase(agenda port.AgendaPort) *AvailabilityUsecase {
	return &AvailabilityUsecase{agenda: agenda}
}

func reservationTypeStringToProto(s string) agendav1.ReservationType {
	switch s {
	case "first_visit":
		return agendav1.ReservationType_RESERVATION_TYPE_FIRST_VISIT
	case "labs":
		return agendav1.ReservationType_RESERVATION_TYPE_LABS
	case "therapy":
		return agendav1.ReservationType_RESERVATION_TYPE_THERAPY
	default:
		return agendav1.ReservationType_RESERVATION_TYPE_FOLLOW_UP
	}
}

func (u *AvailabilityUsecase) GetAvailability(ctx context.Context, doctorID, date, resType string) (*domain.AvailabilityResponse, error) {
	pbType := reservationTypeStringToProto(resType)

	resp, err := u.agenda.GetAvailability(ctx, &agendav1.GetAvailabilityRequest{
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
