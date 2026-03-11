package http

import (
	"net/http"

	"go-challenge-agenda/services/api/internal/domain"
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
	// TODO: this handler returns stub data — wire it to the usecase.
	c.JSON(http.StatusOK, domain.AvailabilityResponse{
		Slots: []domain.AvailableSlot{
			{StartsAt: "2024-03-15T09:00:00Z", EndsAt: "2024-03-15T09:30:00Z"},
			{StartsAt: "2024-03-15T09:30:00Z", EndsAt: "2024-03-15T10:00:00Z"},
			{StartsAt: "2024-03-15T10:00:00Z", EndsAt: "2024-03-15T10:30:00Z"},
		},
		FreeRanges: []domain.TimeRange{
			{From: "2024-03-15T09:00:00Z", To: "2024-03-15T17:00:00Z"},
		},
	})
}
