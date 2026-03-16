# Issue #1 — Wire GET /v1/doctors/:id/availability to real usecase

## What was done

- **File:** `services/api/internal/http/availability.go`
- **Change:** Replaced hardcoded stub response with a call to the availability usecase.

## Implementation

- Read `id` from path (`c.Param("id")`), `date` and `type` from query params.
- Validate: doctor id and date required; `type` optional, default `follow_up`, allowed values `first_visit` or `follow_up`. Return 400 for invalid or missing params.
- Call `h.uc.GetAvailability(ctx, doctorID, date, resType)`; the usecase already calls the agenda gRPC and maps to `domain.AvailabilityResponse`.
- On error, call `c.Error(err)` so the error middleware can map gRPC codes to HTTP (after Issue #4).
- On success, return 200 with the usecase result.
