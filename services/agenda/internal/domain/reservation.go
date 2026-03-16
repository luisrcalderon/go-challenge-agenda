package domain

import "time"

type ReservationType int

const (
	ReservationTypeUnspecified ReservationType = iota
	ReservationTypeFirstVisit
	ReservationTypeFollowUp
	ReservationTypeLabs    // 45 min
	ReservationTypeTherapy // 50 min
)

// SlotDuration returns the duration for a reservation type.
func (t ReservationType) SlotDuration() time.Duration {
	switch t {
	case ReservationTypeFirstVisit:
		return 60 * time.Minute
	case ReservationTypeFollowUp:
		return 30 * time.Minute
	case ReservationTypeLabs:
		return 45 * time.Minute
	case ReservationTypeTherapy:
		return 50 * time.Minute
	default:
		return 30 * time.Minute
	}
}

type ReservationStatus int

const (
	ReservationStatusUnspecified ReservationStatus = iota
	ReservationStatusConfirmed
	ReservationStatusCancelled
)

type Reservation struct {
	ID        string
	DoctorID  string
	PatientID string
	StartsAt  time.Time
	EndsAt    time.Time
	Type      ReservationType
	Status    ReservationStatus
}
