# ADR 0001: API depends on AgendaPort interface instead of gRPC client

## Status

Accepted.

## Context

The API service’s usecases (availability, reservation, user, doctor) need to call the agenda service. The generated gRPC client type (`agendav1.AgendaServiceClient`) was injected directly into usecases and handlers. That led to:

- **Testing:** Hard to unit-test usecases without a real or mock gRPC server.
- **Coupling:** API code depended on generated code and gRPC-specific method signatures (e.g. `...grpc.CallOption`).
- **Flexibility:** Swapping the backend (e.g. REST, in-process) would require changing usecase constructors and call sites.

## Decision

Introduce an **AgendaPort** interface in the API layer that describes only the operations the API needs (GetAvailability, CreateReservation, GetReservation, ListReservations, ListDoctors, GetDoctor, patient and reservation operations). Usecases and handlers depend on this interface, not on `AgendaServiceClient`.

A small **adapter** in the API’s `internal/grpc` package wraps the real gRPC client and implements AgendaPort by delegating each call to the client (without `CallOption`). Main builds the gRPC client, wraps it with `NewAgendaPort(client)`, and passes the port to the router and usecases.

## Consequences

- **Pros:** Usecases and HTTP handlers are testable with a fake implementation of AgendaPort. The API no longer depends on gRPC method signatures. A different backend can be plugged in by implementing the same interface.
- **Cons:** The adapter and the interface must be kept in sync with the methods the API actually uses; new RPCs require a new port method and adapter method.
- **Neutral:** The concrete gRPC client is still created in main and only “dressed” as AgendaPort; the actual I/O remains gRPC.
