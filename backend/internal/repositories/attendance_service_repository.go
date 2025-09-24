package repositories

import (
	db "Backend/internal/models"
	"context"
)

type AttendanceHistoryRepository struct{}

func NewAttendanceHistoryRepository() *AttendanceHistoryRepository {
	return &AttendanceHistoryRepository{}
}

func (attendanceHistoryRepository *AttendanceHistoryRepository) GetDeaconServiceHistory(context context.Context, queries *db.Queries, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetHistoryServiceByDeaconIdRow, error) {
	return queries.GetHistoryServiceByDeaconId(context, deaconPage)
}

func (attendanceHistoryRepository *AttendanceHistoryRepository) AddDeaconServiceHistory(context context.Context, queries *db.Queries, params db.AddHistoryServiceParams) error {
	return queries.AddHistoryService(context, params)
}

func (attendanceHistoryRepository *AttendanceHistoryRepository) DeleteDeaconServiceHistory(context context.Context, queries *db.Queries, params db.DeleteHistoryServiceParams) error {
	return queries.DeleteHistoryService(context, params)
}
