package handlers

import (
	"Backend/internal/middleware"
	"Backend/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeaconHandler struct {
	Service *services.DeaconService
}

func NewDeaconHandler() *DeaconHandler {
	return &DeaconHandler{services.NewDeaconService()}
}

func (deaconHandler *DeaconHandler) GetProfile(c *gin.Context) {
	deaconIdString := c.Param("id")
	if deaconIdString == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty deacon id"})
		return
	}
	deaconId, err := strconv.ParseInt(deaconIdString, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse deacon id: {%s}", deaconIdString)})
		return
	}
	response, err := deaconHandler.Service.GetDeaconProfile(c.Request.Context(), middleware.GetQueries(c), deaconId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
