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
)

type AttendanceHistory struct {
	Service *services.AttendanceHistoryService
}

func NewAttendanceHistoryHandler() *AttendanceHistory {
	return &AttendanceHistory{services.NewAttendanceHistoryServiceService()}
}

// GetServiceHistory @Summary Get All Services done by a deacon
// @Description Get All Services done by a deacon
// @Security BearerAuth
// @Tags Attendance
// @Produce json
// @Param id path int true "Deacon ID"
// @Param page query string true "Page Number"
// @Param limit query string true "Number of elements per page"
// @Success 200 {object} []db.GetHistoryServiceByDeaconIdRow
// @Failure 400
// @Failure 500
// @Router /attendance/deacon/{id} [get]
func (attendanceHistory *AttendanceHistory) GetServiceHistory(c *gin.Context) {
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
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page: {%s}", pageStr)})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page limit {%s}", limitStr)})
		return
	}
	response, err := attendanceHistory.Service.GetServiceHistory(c.Request.Context(), middleware.GetQueries(c), db.GetHistoryServiceByDeaconIdParams{DeaconID: deaconId, Offset: int32((page - 1) * limit), Limit: int32(limit)})
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// AddServiceHistory @Summary Add Service History to a Deacon
// @Description Add Service History to a Deacon
// @Security BearerAuth
// @Tags Attendance
// @Produce json
// @Param id path int true "Deacon ID"
// @Param payload.Attendance body payload.Attendance true "Attendance"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /attendance/deacon/{id} [post]
func (attendanceHistory *AttendanceHistory) AddServiceHistory(c *gin.Context) {
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
	var attendance payload.Attendance
	if err = c.BindJSON(&attendance); err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = attendanceHistory.Service.AddServiceHistory(c.Request.Context(), middleware.GetQueries(c), deaconId, attendance)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

// BulkAssign @Summary Add Service History to a Deacon
// @Description Add Service History to a Deacon
// @Security BearerAuth
// @Tags Attendance
// @Produce json
// @Param payload.BulkAttendance body payload.BulkAttendance true "Attendance"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /attendance/deacon/bulk [post]
func (attendanceHistory *AttendanceHistory) BulkAssign(c *gin.Context) {
	var attendance payload.BulkAttendance
	if err := c.BindJSON(&attendance); err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := attendanceHistory.Service.AddBulkServiceHistory(c.Request.Context(), middleware.GetQueries(c), middleware.GetDatabase(c), attendance)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

// DeleteServiceHistory @Summary Delete Service History to a Deacon
// @Description Delete Service History to a Deacon
// @Security BearerAuth
// @Tags Attendance
// @Produce json
// @Param id path int true "Deacon ID"
// @Param payload.Attendance body payload.Attendance true "Attendance"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /attendance/deacon/{id} [delete]
func (attendanceHistory *AttendanceHistory) DeleteServiceHistory(c *gin.Context) {
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
	var attendance payload.Attendance
	if err = c.BindJSON(&attendance); err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = attendanceHistory.Service.DeleteServiceHistory(c.Request.Context(), middleware.GetQueries(c), deaconId, attendance)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetAllServiceHistory @Summary Get All History
// @Description Get Service History in our church
// @Security BearerAuth
// @Tags Attendance
// @Produce json
// @Param page query string true "Page Number"
// @Param limit query string true "Number of elements per page"
// @Success 200 {object} []db.GetAllHistoryServiceRow
// @Failure 400
// @Failure 500
// @Router /attendance/history [get]
func (attendanceHistory *AttendanceHistory) GetAllServiceHistory(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page: {%s}", pageStr)})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page limit {%s}", limitStr)})
		return
	}
	history, err := attendanceHistory.Service.GetAllServiceHistory(c.Request.Context(), middleware.GetQueries(c), int32((page-1)*limit), int32(limit))
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, history)
}

// GetSuggestion @Summary Get Suggestion
// @Description Get Suggestion for a specific skill
// @Security BearerAuth
// @Tags Attendance
// @Produce json
// @Param eslId query string true "Event Skill Liturgy ID"
// @Param page query string true "Page Number"
// @Param limit query string true "Number of elements per page"
// @Success 200 {object} []db.GetSuggestionRow
// @Failure 400
// @Failure 500
// @Router /attendance/suggestion [get]
func (attendanceHistory *AttendanceHistory) GetSuggestion(c *gin.Context) {
	eventSkillLiturgyIdStr := c.Query("eslId")
	eventSkillLiturgyId, err := strconv.Atoi(eventSkillLiturgyIdStr)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse event ID %s", eventSkillLiturgyIdStr)})
		return
	}
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page: {%s}", pageStr)})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("cannot parse page limit {%s}", limitStr)})
		return
	}
	suggestion, err := attendanceHistory.Service.GetSuggestion(c.Request.Context(), middleware.GetQueries(c), int32(eventSkillLiturgyId), int32((page-1)*limit), int32(limit))
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, suggestion)
}
