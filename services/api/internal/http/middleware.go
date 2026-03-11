package http

import (
	"net/http"

	"go-challenge-agenda/pkg/errcodes"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

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
