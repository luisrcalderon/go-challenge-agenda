# Issue #10 — Route all HTTP through usecase layer

## What was done

Doctor and reservation handlers no longer call the agenda port directly; they use usecases only.

### Doctor usecase

- **File:** `services/api/internal/usecase/doctor.go`
- **DoctorUsecase** with `List(ctx)` and `Get(ctx, id)` that call the agenda port and return `[]domain.DoctorResponse` and `*domain.DoctorResponse`.

### Reservation usecase

- **File:** `services/api/internal/usecase/reservation.go`
- Added **Get(ctx, id)** and **List(ctx, doctorID, from, to)** that call the agenda port and return `*domain.ReservationResponse` and `[]domain.ReservationResponse`.

### Handlers

- **DoctorHandler** now takes only `*usecase.DoctorUsecase`; **List** and **Get** call `uc.List()` and `uc.Get()`. Removed direct port usage and proto conversion from the handler.
- **ReservationHandler** now takes only `*usecase.ReservationUsecase`; **Get** and **List** call `uc.Get()` and `uc.List()`. Removed port and `protoReservationToDTO` from the handler.

### Router

- Creates **DoctorUsecase** and passes it to **NewDoctorHandler(doctorUC)**.
- Passes only **resUC** to **NewReservationHandler(resUC)**.

All HTTP flows now go through the usecase layer; handlers only parse input, call the usecase, and write the response.
