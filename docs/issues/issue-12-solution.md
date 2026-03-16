# Issue #12 — Support additional service types (labs, therapy)

## What was done

Extended reservation types with **labs** (45 min) and **therapy** (50 min); domain, proto, and API now support them with no change to business logic structure.

### Proto

- **File:** `proto/agenda/v1/agenda.proto`
- Added `RESERVATION_TYPE_LABS = 3` and `RESERVATION_TYPE_THERAPY = 4` to `ReservationType` enum. Regenerated with `make proto`.

### Domain

- **File:** `services/agenda/internal/domain/reservation.go`
- Added `ReservationTypeLabs` and `ReservationTypeTherapy`; **SlotDuration()** returns 45m for labs, 50m for therapy. Existing usecases (availability, reservation) already use `SlotDuration()` and the type enum, so they pick up the new types.

### API

- **Availability:** Handler allows `type=labs` and `type=therapy`; usecase maps string to proto via **reservationTypeStringToProto** (in `usecase/availability.go`).
- **Reservation:** Create accepts `type` `labs`/`therapy`; **reservationTypeStringToProto** used when calling agenda. Response DTOs map proto back to string via **reservationTypeProtoToString** (in `usecase/reservation.go`), used in **protoReservationToDTO** and in user usecase **reservationProtoToDTO**.

### Test doubles

- **availability_bench_test.go** and **reservation_race_test.go:** Implemented **ListReservationsByPatient** on fake reservation repos so they satisfy the updated `ReservationRepository` interface (required after Issue #6).

No repository schema change: reservation type is already stored as an int; new enum values are stored as 3 and 4.
