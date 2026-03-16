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
// Expands daily, weekly, or monthly recurrence until RecurrenceUntil (or until past to).
func (b *BlockedSlot) Occurrences(from, to time.Time) []BlockedSlot {
	duration := b.EndsAt.Sub(b.StartsAt)

	switch b.RecurrenceType {
	case RecurrenceNone:
		if b.StartsAt.After(to) || b.EndsAt.Before(from) {
			return nil
		}
		return []BlockedSlot{*b}
	case RecurrenceDaily:
		return b.occurrencesByStep(from, to, duration, func(t time.Time) time.Time { return t.AddDate(0, 0, 1) })
	case RecurrenceWeekly:
		return b.occurrencesByStep(from, to, duration, func(t time.Time) time.Time { return t.AddDate(0, 0, 7) })
	case RecurrenceMonthly:
		return b.occurrencesByStep(from, to, duration, func(t time.Time) time.Time { return t.AddDate(0, 1, 0) })
	default:
		if b.StartsAt.After(to) || b.EndsAt.Before(from) {
			return nil
		}
		return []BlockedSlot{*b}
	}
}

// occurrencesByStep generates occurrences by advancing start with step until RecurrenceUntil or past to.
func (b *BlockedSlot) occurrencesByStep(from, to time.Time, duration time.Duration, step func(time.Time) time.Time) []BlockedSlot {
	var out []BlockedSlot
	start := b.StartsAt
	for {
		if start.After(to) {
			break
		}
		if b.RecurrenceUntil != nil && start.After(*b.RecurrenceUntil) {
			break
		}
		end := start.Add(duration)
		if end.After(from) && start.Before(to) {
			out = append(out, BlockedSlot{
				ID:             b.ID,
				DoctorID:       b.DoctorID,
				StartsAt:       start,
				EndsAt:         end,
				Reason:         b.Reason,
				RecurrenceType: RecurrenceNone,
				RecurrenceUntil: nil,
			})
		}
		start = step(start)
	}
	return out
}
