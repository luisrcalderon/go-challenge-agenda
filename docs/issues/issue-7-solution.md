# Issue #7 — Decouple API usecases from concrete gRPC client

## What was done

Introduced an `AgendaPort` interface and an adapter so API usecases and HTTP handlers depend on the port instead of the generated gRPC client type.

### New: port package

- **File:** `services/api/internal/port/agenda.go`
- **AgendaPort** interface with the methods the API needs: GetAvailability, CreateReservation, CancelReservation, GetReservation, ListReservations, ListDoctors, GetDoctor, ListPatients, GetPatient, CreatePatient, UpdatePatient, DeletePatient. Signatures use proto request/response types but no `...grpc.CallOption`, so the interface is independent of gRPC.

### Adapter

- **File:** `services/api/internal/grpc/client.go`
- **agendaPortAdapter** wraps `agendav1.AgendaServiceClient` and implements `port.AgendaPort` by delegating each call to the client (no call options). **NewAgendaPort(client)** returns `port.AgendaPort`.

### Usecases

- **AvailabilityUsecase,** **ReservationUsecase,** **UserUsecase** now take `port.AgendaPort` instead of `agendav1.AgendaServiceClient`; internal calls use the port.

### Handlers and router

- **DoctorHandler** and **ReservationHandler** take `port.AgendaPort` instead of the gRPC client.
- **NewRouter(agenda port.AgendaPort)**; main builds the client, then `agenda := NewAgendaPort(client)` and passes `agenda` to NewRouter.

Tests can inject a fake that implements AgendaPort without using gRPC.
