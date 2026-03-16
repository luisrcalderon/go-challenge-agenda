# Proposal: Dedicated notifications service (Issue #14)

## Benefit

A separate **notifications** service would:

- Send reminders (e.g. email/SMS) before appointments.
- Notify doctors or staff when a reservation is created or cancelled.
- Keep notification logic and delivery (templates, retries, providers) out of the agenda and API services.

Today, the API and agenda services have no notification behaviour; adding it inside either would mix scheduling with delivery concerns and make scaling/ownership harder.

## Service boundary

- **Owns:** Sending messages (email, SMS, push), template rendering, retry and backoff, provider integration.
- **Does not own:** Scheduling, availability, or user/doctor data; it consumes events or explicit “send” calls from other services.

## Proto contract (sketch)

```protobuf
syntax = "proto3";
package notifications.v1;

service NotificationsService {
  // Send a single notification (e.g. reservation reminder).
  rpc Send(SendRequest) returns (SendResponse);
  // Optional: subscribe to reservation events from agenda (event-driven).
  // rpc OnReservationCreated(ReservationCreatedEvent) returns (Ack);
}

message SendRequest {
  string channel = 1;  // "email" | "sms"
  string recipient = 2;
  string template_id = 3;
  map<string, string> template_data = 4;
}

message SendResponse {
  string id = 1;
  bool accepted = 2;
}
```

## Integration

- **Option A (call from API):** After creating a reservation via the agenda client, the API calls `NotificationsService.Send` with a “reservation_confirmation” template and recipient from the request.
- **Option B (event-driven):** Agenda service publishes “ReservationCreated” to a queue; the notifications service subscribes and sends reminders. Requires a message bus and keeps agenda free of notification code.

## Implementation scope (optional)

A minimal implementation could be a third binary `services/notifications` that implements the above proto and, for development, logs or writes notifications to a file instead of calling real providers. The API would then call this service after a successful reservation create (Option A) when configured.
