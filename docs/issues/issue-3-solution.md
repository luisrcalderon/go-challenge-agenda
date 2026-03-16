# Issue #3 — Implement ListReservations in SQLite repository

## What was done

- **File:** `services/agenda/internal/repository/sqlite/reservation.go`
- **Change:** Implemented `ListReservations` instead of returning `(nil, nil)`.

## Implementation

- Query the `reservations` table with:
  - `doctor_id = ?`
  - Overlap condition: `starts_at < to AND ends_at > from` (a reservation overlaps `[from, to]` when it starts before `to` and ends after `from`).
- Map each row to `*domain.Reservation` via `models.ReservationFromModel`.
- Return the slice and any database error.

This unblocks real availability data (Issue #1), the list-by-patient pattern (Issue #6), and TestListReservations (Issue #5).
