# Issue #4 — Map gRPC errors to HTTP status codes

## What was done

- **File:** `pkg/errcodes/mapping.go`
- **Change:** Extended `GRPCToHTTP` so more gRPC codes map to appropriate HTTP status codes instead of falling through to 500.

## Mappings added

| gRPC Code           | HTTP Status |
|---------------------|------------|
| AlreadyExists       | 409 Conflict |
| FailedPrecondition  | 409 Conflict |
| PermissionDenied    | 403 Forbidden |
| Unauthenticated     | 401 Unauthorized |
| ResourceExhausted   | 429 Too Many Requests |
| Unimplemented       | 501 Not Implemented |
| DeadlineExceeded   | 504 Gateway Timeout |
| Canceled            | 408 Request Timeout |
| Aborted             | 409 Conflict |

Conflict (reservation “time slot not available”) is now returned as 409 thanks to the agenda gRPC returning `FailedPrecondition` (Issue #2) and this mapping.
