package handlers

import (
	"Backend/internal/middleware"
	"Backend/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
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
