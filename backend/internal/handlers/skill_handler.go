package handlers

import (
	"Backend/internal/middleware"
	db "Backend/internal/models"
	"Backend/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SkillHandler struct {
	Service *services.SkillService
}

func NewSkillHandler() *SkillHandler {
	return &SkillHandler{services.NewSkillService()}
}

// GetSkills @Summary Lists All Skills
// @Description Lists All Skills
// @Security BearerAuth
// @Tags Skills
// @Produce json
// @Success 200 {object} db.GetSkillsRow
// @Failure 400
// @Failure 500
// @Router /skills [get]
func (skillHandler *SkillHandler) GetSkills(c *gin.Context) {
	response, err := skillHandler.Service.GetSkills(c.Request.Context(), middleware.GetQueries(c))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetDependantSkills @Summary Lists All Depending Skills
// @Description Lists All Skills depending on liturgy and event
// @Security BearerAuth
// @Tags Skills
// @Produce json
// @Param liturgy_id path int true "Liturgy ID"
// @Param event_id path int true "Event ID"
// @Success 200 {object} db.GetDependantSkillsRow
// @Failure 400
// @Failure 500
// @Router /skills/{liturgy_id}/{event_id} [get]
func (skillHandler *SkillHandler) GetDependantSkills(c *gin.Context) {
	liturgyIdString := c.Param("liturgy_id")
	if liturgyIdString == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty liturgy id"})
		return
	}
	liturgyId, err := strconv.ParseInt(liturgyIdString, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse liturgy id: {%s}", liturgyIdString)})
		return
	}
	eventIdString := c.Param("event_id")
	if eventIdString == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty event id"})
		return
	}
	eventId, err := strconv.ParseInt(eventIdString, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse event id: {%s}", eventIdString)})
		return
	}
	response, err := skillHandler.Service.GetDependantSkills(c.Request.Context(), middleware.GetQueries(c), db.GetDependantSkillsParams{LiturgyID: int32(liturgyId), EventID: int32(eventId)})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
