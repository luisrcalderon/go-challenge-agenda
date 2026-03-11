package errcodes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// GRPCToHTTP maps gRPC status codes to HTTP status codes.
// TODO: mapping is incomplete — most codes fall through to 500.
func GRPCToHTTP(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.NotFound:
		return http.StatusNotFound
	case codes.InvalidArgument:
		return http.StatusBadRequest
	// TODO: add mappings for AlreadyExists, PermissionDenied, Unauthenticated, ResourceExhausted, etc.
	default:
		return http.StatusInternalServerError
	}
}
