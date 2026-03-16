# Issue #2 — Fix reservation conflict detection and slot duration

## What was done

### 1. Conflict detection (`hasConflict`)

- **File:** `services/agenda/internal/usecase/reservation.go`
- **Problem:** Only checked whether the new start time fell inside an existing slot. Missed (1) new slot fully containing an existing one, (2) new slot ending inside an existing one.
- **Change:** Use proper interval overlap: two ranges overlap when `startsAt.Before(r.EndsAt) && endsAt.After(r.StartsAt)`. Kept the filter that skips cancelled reservations.

### 2. Slot duration

- **File:** `services/agenda/internal/domain/reservation.go`
- **Problem:** `SlotDuration()` always returned 30 minutes; first visits should be 60 minutes.
- **Change:** Return 60m for `ReservationTypeFirstVisit`, 30m for `ReservationTypeFollowUp`, 30m for default/unspecified.

### 3. Conflict as gRPC status (for Issue #4 mapping)

- **Files:** `services/agenda/internal/usecase/reservation.go`, `services/agenda/internal/grpc/server.go`
- **Change:** Introduced `usecase.ErrConflict` and have `Create` return it on conflict. In the gRPC handler, detect `ErrConflict` and return `codes.FailedPrecondition` with "time slot not available" so the API can map to 409.
