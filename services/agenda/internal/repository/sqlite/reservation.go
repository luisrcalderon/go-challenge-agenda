package sqlite

import (
	"context"
	"fmt"
	"time"

	"go-challenge-agenda/services/agenda/internal/domain"
	"go-challenge-agenda/services/agenda/internal/repository/models"

	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) CreateReservation(ctx context.Context, res *domain.Reservation) error {
	m := models.ReservationToModel(res)
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *ReservationRepository) GetReservation(ctx context.Context, id string) (*domain.Reservation, error) {
	var m models.Reservation
	res := r.db.WithContext(ctx).First(&m, "id = ?", id)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("reservation not found: %s", id)
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return models.ReservationFromModel(&m), nil
}

// ListReservations returns reservations for a doctor overlapping [from, to].
// TODO: implement.
func (r *ReservationRepository) ListReservations(ctx context.Context, doctorID string, from, to time.Time) ([]*domain.Reservation, error) {
	return nil, nil
}

func (r *ReservationRepository) UpdateReservation(ctx context.Context, res *domain.Reservation) error {
	m := models.ReservationToModel(res)
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *ReservationRepository) CancelReservation(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).
		Model(&models.Reservation{}).
		Where("id = ?", id).
		Update("status", int(domain.ReservationStatusCancelled)).Error
}
