package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"go-challenge-agenda/pkg/errcodes"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc/status"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests by method and status",
		},
		[]string{"method", "status"},
	)
)

// LoggingMiddleware logs each request with method, path, status, and duration using structured logging.
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		c.Next()
		status := c.Writer.Status()
		duration := time.Since(start)
		slog.Info("request",
			"method", method,
			"path", path,
			"status", status,
			"duration_ms", duration.Milliseconds(),
		)
		httpRequestsTotal.WithLabelValues(method, fmt.Sprintf("%d", status)).Inc()
	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}
		err := c.Errors.Last().Err
		st, ok := status.FromError(err)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		httpCode := errcodes.GRPCToHTTP(st.Code())
		c.JSON(httpCode, gin.H{"error": st.Message()})
	}
}
