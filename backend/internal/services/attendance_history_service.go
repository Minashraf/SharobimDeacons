package services

import (
	"Backend/internal/data/payload"
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"context"
	"time"
)

type AttendanceHistoryService struct {
	Repository *repositories.AttendanceHistoryRepository
}

func NewAttendanceHistoryServiceService() *AttendanceHistoryService {
	return &AttendanceHistoryService{Repository: repositories.NewAttendanceHistoryRepository()}
}

func (s AttendanceHistoryService) GetServiceHistory(context context.Context, queries *db.Queries, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetHistoryServiceByDeaconIdRow, error) {
	return s.Repository.GetDeaconServiceHistory(context, queries, deaconPage)
}

func (s AttendanceHistoryService) AddServiceHistory(context context.Context, queries *db.Queries, deaconId int64, attendance payload.Attendance) error {
	dateOnly, err := time.Parse(time.DateOnly, attendance.Date)
	if err != nil {
		return err
	}
	return s.Repository.AddDeaconServiceHistory(context, queries, db.AddHistoryServiceParams{DeaconID: deaconId, Date: dateOnly, EventSkillLiturgyID: attendance.ESLId})
}

func (s AttendanceHistoryService) DeleteServiceHistory(context context.Context, queries *db.Queries, deaconId int64, attendance payload.Attendance) error {
	dateOnly, err := time.Parse(time.DateOnly, attendance.Date)
	if err != nil {
		return err
	}
	return s.Repository.DeleteDeaconServiceHistory(context, queries, db.DeleteHistoryServiceParams{DeaconID: deaconId, Date: dateOnly, EventSkillLiturgyID: attendance.ESLId})
}
