# Issue #13 — Structured logging and at least one metric

## What was done

Added observability at the transport layer only: structured logging (slog) and a Prometheus HTTP request counter. No changes in domain or usecase layers.

### API service (HTTP)

- **File:** `services/api/internal/http/middleware.go`
  - **LoggingMiddleware:** Runs after the handler; logs each request with `slog.Info("request", "method", method, "path", path, "status", status, "duration_ms", duration.Milliseconds())`.
  - **Metric:** `http_requests_total` counter (Prometheus) with labels `method` and `status` (numeric), incremented in the same middleware.
- **File:** `services/api/internal/http/router.go`
  - Router uses `LoggingMiddleware()` before `ErrorMiddleware()`.
  - `GET /metrics` registered with `promhttp.Handler()` to expose Prometheus metrics.

### Agenda service (gRPC)

- **File:** `services/agenda/cmd/main.go`
  - **grpcLoggingInterceptor:** Unary server interceptor that logs after each RPC: `slog.Info("grpc_request", "method", info.FullMethod, "code", code.String(), "duration_ms", duration.Milliseconds())`. Code is taken from the gRPC status when the handler returns an error.
  - Server created with `grpc.NewServer(grpc.UnaryInterceptor(grpcLoggingInterceptor))`.

### Dependencies

- **go.mod:** `github.com/prometheus/client_golang v1.20.5` for the counter and `promhttp.Handler()`. Standard library `log/slog` for logging.

Logging and metrics are confined to the edge (HTTP middleware, gRPC interceptor); domain and usecases stay free of observability code.
