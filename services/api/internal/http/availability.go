package http

import (
	"net/http"

	"go-challenge-agenda/services/api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AvailabilityHandler struct {
	uc *usecase.AvailabilityUsecase
}

func NewAvailabilityHandler(uc *usecase.AvailabilityUsecase) *AvailabilityHandler {
	return &AvailabilityHandler{uc: uc}
}

// Get godoc
// @Summary     Get available slots for a doctor on a given date
// @Tags        availability
// @Produce     json
// @Param       id    path      string  true   "Doctor ID"
// @Param       date  query     string  true   "Date (YYYY-MM-DD)"
// @Param       type  query     string  false  "Reservation type: first_visit or follow_up"
// @Success     200   {object}  domain.AvailabilityResponse
// @Failure     400   {object}  map[string]string
// @Failure     500   {object}  map[string]string
// @Router      /doctors/{id}/availability [get]
func (h *AvailabilityHandler) Get(c *gin.Context) {
	doctorID := c.Param("id")
	if doctorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doctor id is required"})
		return
	}
	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date query param is required (YYYY-MM-DD)"})
		return
	}
	resType := c.DefaultQuery("type", "follow_up")
	allowed := map[string]bool{"first_visit": true, "follow_up": true, "labs": true, "therapy": true}
	if !allowed[resType] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type must be first_visit, follow_up, labs, or therapy"})
		return
	}

	result, err := h.uc.GetAvailability(c.Request.Context(), doctorID, date, resType)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, result)
}
