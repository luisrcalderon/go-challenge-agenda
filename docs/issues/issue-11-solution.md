# Issue #11 — Include blocked slots in availability

## What was done

- **File:** `services/agenda/internal/usecase/availability.go`
- **Change:** Fetched blocked slots for the doctor and day, expanded recurrences, and merged them into the busy set before computing free ranges and slots.

## Implementation

- Call `u.blockedSlots.ListBlockedSlots(ctx, doctorID, dayStart, dayEnd)` to get stored blocked slots overlapping the working window.
- For each slot, call `Occurrences(dayStart, dayEnd)` (from Issue #8) to expand daily/weekly/monthly recurrence.
- Append each occurrence’s `[StartsAt, EndsAt]` to the `busy` slice (same shape as reservation busy ranges).
- Reuse existing `subtractBusy(dayStart, dayEnd, busy)` so blocked time is subtracted from free time together with reservations; slots are then derived from the resulting free ranges.

Availability no longer shows slots that fall in blocked periods.
