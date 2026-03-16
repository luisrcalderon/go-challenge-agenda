# Issue #6 — Implement GET /v1/users/:id/reservations end-to-end

## What was done

Implemented listing reservations by user (patient) from the API down to the repository.

### Proto

- **File:** `proto/agenda/v1/agenda.proto`
- Added optional `patient_id` (field 4) to `ListReservationsRequest`. When set, list reservations for that patient; when `doctor_id` is set, keep existing behavior. Regenerated with `make proto`.

### Agenda domain

- **File:** `services/agenda/internal/domain/repository.go`
- Added `ListReservationsByPatient(ctx, patientID, from, to)` to `ReservationRepository`.

### Agenda SQLite repository

- **File:** `services/agenda/internal/repository/sqlite/reservation.go`
- Implemented `ListReservationsByPatient`: query by `patient_id` with same overlap condition (`starts_at < to AND ends_at > from`), map to `[]*domain.Reservation`.

### Agenda usecase

- **File:** `services/agenda/internal/usecase/reservation.go`
- Added `ListByPatient(ctx, patientID, from, to)` calling the new repo method.

### Agenda gRPC server

- **File:** `services/agenda/internal/grpc/server.go`
- `ListReservations`: if `req.From`/`req.To` are empty, use default range (now −1 year to now +1 year). If `req.PatientId != ""`, call `ListByPatient`; otherwise call `List` with `DoctorId`.

### API user usecase

- **File:** `services/api/internal/usecase/user.go`
- Added `ListReservations(ctx, userID)` calling agenda `ListReservations` with `PatientId: userID` (no from/to, so agenda uses default range). Map response to `[]domain.ReservationResponse` via `reservationProtoToDTO`.

### API user handler

- **File:** `services/api/internal/http/user.go`
- `ListReservations`: validate user id, call `uc.Get` to ensure user exists (404 if not), then `uc.ListReservations(userID)`, return 200 with the list.

### Mock

- **File:** `services/agenda/internal/domain/mocks/ReservationRepository.go`
- Added `ListReservationsByPatient` to the mock so the interface is satisfied.
