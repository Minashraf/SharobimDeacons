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

func (deaconHandler *DeaconHandler) GetServiceHistory(c *gin.Context) {
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
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot page: {%s}", pageStr)})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page limit {%s}", limitStr)})
		return
	}
	response, err := deaconHandler.Service.GetServiceHistory(c.Request.Context(), middleware.GetQueries(c), db.GetHistoryServiceByDeaconIdParams{DeaconID: deaconId, Offset: int32(page) - 1, Limit: int32(limit)})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (deaconHandler *DeaconHandler) GetDeacons(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	sortField := c.DefaultQuery("sort", "first_name")
	sortDirection := c.DefaultQuery("direction", "asc")
	sorting := make(map[string]string)
	sorting["Field"] = sortField
	sorting["Direction"] = sortDirection
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot page: {%s}", pageStr)})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page limit {%s}", limitStr)})
		return
	}
	response, err := deaconHandler.Service.GetDeacons(c.Request.Context(), middleware.GetQueries(c), sorting, db.GetHistoryServiceByDeaconIdParams{Offset: int32(page - 1), Limit: int32(limit)})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
