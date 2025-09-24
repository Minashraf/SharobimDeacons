package handlers

import (
	"Backend/internal/middleware"
	"Backend/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LiturgyHandler struct {
	Service *services.LiturgyService
}

func NewLiturgyHandler() *LiturgyHandler {
	return &LiturgyHandler{services.NewLiturgyService()}
}

// GetLiturgies @Summary Lists All Liturgies
// @Description Lists All Liturgies
// @Security BearerAuth
// @Tags Liturgies
// @Produce json
// @Success 200 {object} db.GetLiturgiesRow
// @Failure 400
// @Failure 500
// @Router /liturgy [get]
func (liturgyHandler *LiturgyHandler) GetLiturgies(c *gin.Context) {
	response, err := liturgyHandler.Service.GetLiturgies(c.Request.Context(), middleware.GetQueries(c))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
