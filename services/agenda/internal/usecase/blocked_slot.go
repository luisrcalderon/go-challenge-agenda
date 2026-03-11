package usecase

import (
	"context"
	"fmt"
	"time"

	"go-challenge-agenda/services/agenda/internal/domain"

	"github.com/google/uuid"
)

type BlockedSlotUsecase struct {
	repo domain.BlockedSlotRepository
}

func NewBlockedSlotUsecase(repo domain.BlockedSlotRepository) *BlockedSlotUsecase {
	return &BlockedSlotUsecase{repo: repo}
}

func (u *BlockedSlotUsecase) Create(ctx context.Context, b *domain.BlockedSlot) (*domain.BlockedSlot, error) {
	b.ID = uuid.NewString()
	if err := u.repo.CreateBlockedSlot(ctx, b); err != nil {
		return nil, fmt.Errorf("create blocked slot: %w", err)
	}
	return b, nil
}

func (u *BlockedSlotUsecase) List(ctx context.Context, doctorID string, from, to time.Time) ([]*domain.BlockedSlot, error) {
	base, err := u.repo.ListBlockedSlots(ctx, doctorID, from, to)
	if err != nil {
		return nil, err
	}

	// Expand recurrences — each slot's Occurrences() handles expansion.

	var expanded []*domain.BlockedSlot
	for _, b := range base {
		for _, occ := range b.Occurrences(from, to) {
			occ := occ
			expanded = append(expanded, &occ)
		}
	}
	return expanded, nil
}

func (u *BlockedSlotUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.DeleteBlockedSlot(ctx, id)
}
