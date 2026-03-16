# Issue #8 — Implement recurrence expansion for blocked slots

## What was done

- **File:** `services/agenda/internal/domain/blocked_slot.go`
- **Change:** Implemented `Occurrences(from, to)` to expand daily, weekly, and monthly recurrence.

## Implementation

- **RecurrenceNone:** Return the single slot if it overlaps `[from, to]`; otherwise nil.
- **RecurrenceDaily:** From `StartsAt`, advance by 24h (AddDate(0, 0, 1)) until after `to` or after `RecurrenceUntil`. Each occurrence is a copy with shifted `StartsAt`/`EndsAt` (same duration), included only if it overlaps `[from, to]`.
- **RecurrenceWeekly:** Same logic with 7-day step (AddDate(0, 0, 7)).
- **RecurrenceMonthly:** Same logic with 1-month step (AddDate(0, 1, 0)).
- Added helper `occurrencesByStep(from, to, duration, step)` to avoid duplication; each occurrence is emitted with `RecurrenceType: RecurrenceNone` and `RecurrenceUntil: nil` so downstream logic sees concrete time ranges.

This allows the availability usecase (Issue #11) to merge expanded blocked ranges into busy time.
