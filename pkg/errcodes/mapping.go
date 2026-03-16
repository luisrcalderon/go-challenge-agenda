package errcodes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// GRPCToHTTP maps gRPC status codes to HTTP status codes.
func GRPCToHTTP(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.NotFound:
		return http.StatusNotFound
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.FailedPrecondition:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Aborted:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
