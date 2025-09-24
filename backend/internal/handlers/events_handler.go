package handlers

import (
	"Backend/internal/middleware"
	"Backend/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EventsHandler struct {
	Service *services.EventsService
}

func NewEventsHandler() *EventsHandler {
	return &EventsHandler{services.NewEventsService()}
}

// GetEvents @Summary Lists All Skills
// @Description Lists All Skills
// @Security BearerAuth
// @Tags Events
// @Param liturgy_id path int true "Liturgy ID"
// @Produce json
// @Success 200 {object} db.GetEventsRow
// @Failure 400
// @Failure 500
// @Router /events/{liturgy_id} [get]
func (eventsHandler *EventsHandler) GetEvents(c *gin.Context) {
	liturgyIdString := c.Param("liturgy_id")
	if liturgyIdString == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty liturgy id"})
		return
	}
	liturgyId, err := strconv.ParseInt(liturgyIdString, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse liturgy id: {%s}", liturgyId)})
		return
	}
	response, err := eventsHandler.Service.GetEvents(c.Request.Context(), middleware.GetQueries(c), int32(liturgyId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
