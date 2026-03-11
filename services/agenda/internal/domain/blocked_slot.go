package domain

import "time"

type RecurrenceType int

const (
	RecurrenceNone RecurrenceType = iota
	RecurrenceDaily
	RecurrenceWeekly
	RecurrenceMonthly
)

type BlockedSlot struct {
	ID              string
	DoctorID        string
	StartsAt        time.Time
	EndsAt          time.Time
	Reason          string
	RecurrenceType  RecurrenceType
	RecurrenceUntil *time.Time
}

// Occurrences returns all occurrences of this blocked slot within [from, to].
// TODO: recurrence expansion is not implemented — only returns the base slot.
func (b *BlockedSlot) Occurrences(from, to time.Time) []BlockedSlot {
	if b.StartsAt.After(to) || b.EndsAt.Before(from) {
		return nil
	}
	return []BlockedSlot{*b}
}
