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
// A reservation overlaps when starts_at < to AND ends_at > from.
func (r *ReservationRepository) ListReservations(ctx context.Context, doctorID string, from, to time.Time) ([]*domain.Reservation, error) {
	var ms []models.Reservation
	err := r.db.WithContext(ctx).
		Where("doctor_id = ? AND starts_at < ? AND ends_at > ?", doctorID, to.UTC(), from.UTC()).
		Find(&ms).Error
	if err != nil {
		return nil, err
	}
	out := make([]*domain.Reservation, len(ms))
	for i := range ms {
		out[i] = models.ReservationFromModel(&ms[i])
	}
	return out, nil
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
