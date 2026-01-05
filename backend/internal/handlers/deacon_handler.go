package handlers

import (
	"Backend/internal/data/payload"
	"Backend/internal/middleware"
	db "Backend/internal/models"
	"Backend/internal/services"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type DeaconHandler struct {
	Service *services.DeaconService
}

func NewDeaconHandler() *DeaconHandler {
	return &DeaconHandler{services.NewDeaconService()}
}

// GetDeacon @Summary Get Info of a Deacon
// @Description Get A Detailed info of a specific deacon
// @Security BearerAuth
// @Tags Deacons
// @Produce json
// @Param id path int true "Deacon ID"
// @Success 200 {object} map[string]string
// @Failure 400
// @Failure 500
// @Router /deacons/{id} [get]
func (deaconHandler *DeaconHandler) GetDeacon(c *gin.Context) {
	deaconIdString := c.Param("id")
	if deaconIdString == "" {
		err := errors.New("empty deacon id")
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deaconId, err := strconv.ParseInt(deaconIdString, 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse deacon id: {%s}", deaconIdString)})
		return
	}
	response, err := deaconHandler.Service.GetDeaconProfile(c.Request.Context(), middleware.GetQueries(c), deaconId)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetDeacons @Summary Lists All Deacons
// @Description Lists All Deacons Paginated
// @Security BearerAuth
// @Tags Deacons
// @Produce json
// @Param page query string false "Page Number"
// @Param limit query string false "Number of elements per page"
// @Param sort query string false "name, date_of_birth, country, deacon_rank_id"
// @Param direction query string false "ASC OR DESC"
// @Param filter_field query string false "name, country, deacon_rank_id"
// @Param filter_value query string false "The value of the filtered selection"
// @Success 200 {object} []map[string]string
// @Failure 400
// @Failure 500
// @Router /deacons/ [get]
func (deaconHandler *DeaconHandler) GetDeacons(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	sortingAndFilter := make(map[string]string)
	sortingAndFilter["Field"] = c.DefaultQuery("sort", "first_name")
	sortingAndFilter["Direction"] = strings.ToLower(c.DefaultQuery("direction", "asc"))
	sortingAndFilter["FilterField"] = c.DefaultQuery("filter_field", "")
	sortingAndFilter["FilterValue"] = c.DefaultQuery("filter_value", "")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page: {%s}", pageStr)})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page limit {%s}", limitStr)})
		return
	}
	response, err := deaconHandler.Service.GetDeacons(c.Request.Context(), middleware.GetQueries(c), sortingAndFilter, db.GetHistoryServiceByDeaconIdParams{Offset: int32((page - 1) * limit), Limit: int32(limit)})
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetDeaconsRanks @Summary Lists All Deacons Ranks
// @Description Lists All Deacons Ranks
// @Security BearerAuth
// @Tags Deacons
// @Produce json
// @Success 200 {object} []map[string]string
// @Failure 400
// @Failure 500
// @Router /deacons/ranks [get]
func (deaconHandler *DeaconHandler) GetDeaconsRanks(c *gin.Context) {
	response, err := deaconHandler.Service.GetDeaconsRanks(c.Request.Context(), middleware.GetQueries(c))
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateDeacon @Summary Creates a deacon
// @Description Create a deacon profile
// @Security BearerAuth
// @Tags Deacons
// @Produce json
// @Param payload.Deacon body payload.Deacon true "Creation"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /deacons/ [post]
func (deaconHandler *DeaconHandler) CreateDeacon(c *gin.Context) {
	var deacon payload.Deacon
	if err := c.BindJSON(&deacon); err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := deaconHandler.Service.AddDeacon(c.Request.Context(), middleware.GetQueries(c), middleware.GetDatabase(c), deacon)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

// DeleteDeacon @Summary Deletes a Deacon
// @Description Deletes a Deacon and his skills
// @Security BearerAuth
// @Tags Deacons
// @Produce json
// @Param id path int true "Deacon ID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /deacons/{id} [delete]
func (deaconHandler *DeaconHandler) DeleteDeacon(c *gin.Context) {
	deaconIdString := c.Param("id")
	if deaconIdString == "" {
		err := errors.New("empty deacon id")
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deaconId, err := strconv.ParseInt(deaconIdString, 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse deacon id: {%s}", deaconIdString)})
		return
	}
	err = deaconHandler.Service.DeleteDeacon(c.Request.Context(), middleware.GetQueries(c), middleware.GetDatabase(c), deaconId)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateDeacon @Summary Update a deacon
// @Description Update a deacon and its skills
// @Security BearerAuth
// @Tags Deacons
// @Produce json
// @Param id path int true "Deacon ID"
// @Param payload.Deacon body payload.Deacon true "Creation"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /deacons/{id} [put]
func (deaconHandler *DeaconHandler) UpdateDeacon(c *gin.Context) {
	deaconIdString := c.Param("id")
	if deaconIdString == "" {
		err := errors.New("empty deacon id")
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deaconId, err := strconv.ParseInt(deaconIdString, 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse deacon id: {%s}", deaconIdString)})
		return
	}
	var deacon payload.Deacon
	if err = c.BindJSON(&deacon); err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = deaconHandler.Service.UpdateDeacon(c.Request.Context(), middleware.GetQueries(c), middleware.GetDatabase(c), deacon, deaconId)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
