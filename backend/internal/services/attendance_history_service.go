package services

import (
	"Backend/internal/data/payload"
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"context"
	"database/sql"
	"errors"
	"time"
)

type AttendanceHistoryService struct {
	HistoryRepository *repositories.AttendanceHistoryRepository
	DeaconRepository  *repositories.DeaconRepository
	ESLRepository     *repositories.EventSkillLiturgyRepository
}

var ErrCapacityExceeded = errors.New("you have exceeded the maximum amount of the event capacity")

func NewAttendanceHistoryServiceService() *AttendanceHistoryService {
	return &AttendanceHistoryService{HistoryRepository: repositories.NewAttendanceHistoryRepository(), DeaconRepository: repositories.NewDeaconRepository(), ESLRepository: repositories.NewEventSkillLiturgyRepository()}
}

func (s AttendanceHistoryService) GetServiceHistory(context context.Context, queries *db.Queries, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetHistoryServiceByDeaconIdRow, error) {
	return s.HistoryRepository.GetDeaconServiceHistory(context, queries, deaconPage)
}

func (s AttendanceHistoryService) addServiceHistory(context context.Context, queries *db.Queries, deaconId int64, attendance payload.Attendance, createdBy sql.NullInt64) error {
	dateOnly, err := time.Parse(time.DateOnly, attendance.Date)
	if err != nil {
		return err
	}
	skills, err := s.DeaconRepository.GetDeaconSkills(context, queries, deaconId)
	if err != nil {
		return err
	}
	esl, err := s.ESLRepository.GetEventSkillLiturgy(context, queries, attendance.ESLId)
	if err != nil {
		return err
	}
	for _, skill := range skills {
		if esl.SkillID == skill.ID {
			return s.HistoryRepository.AddDeaconServiceHistory(context, queries, db.AddHistoryServiceParams{DeaconID: deaconId, Date: dateOnly, EventSkillLiturgyID: attendance.ESLId, CreatedBy: createdBy})
		}
	}
	return errors.New("the Deacon doesn't have the required Skill")
}

func (s AttendanceHistoryService) AddBulkServiceHistory(context context.Context, queries *db.Queries, database *sql.DB, attendance payload.BulkAttendance) error {
	dateOnly, err := time.Parse(time.DateOnly, attendance.Date)
	if err != nil {
		return err
	}

	remaining, err := s.ESLRepository.GetRemainingEventSkillLiturgyCapacity(context, queries, db.GetRemainingEventSkillLiturgyCapacityParams{Eslid: attendance.ESLId, ServiceDate: dateOnly})
	if err != nil {
		return err
	}
	if !attendance.OverrideWarning && int(remaining) < len(attendance.DeaconId) {
		return ErrCapacityExceeded
	}
	tx, err := database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := queries.WithTx(tx)

	userId, ok := context.Value("UserId").(int64)

	var createdBy sql.NullInt64
	if ok {
		createdBy = sql.NullInt64{
			Int64: userId,
			Valid: true,
		}
	} else {
		return errors.New("UserId is not int64")
	}

	for _, deaconId := range attendance.DeaconId {
		err = s.addServiceHistory(context, qtx, deaconId, payload.Attendance{ESLId: attendance.ESLId, Date: attendance.Date}, createdBy)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (s AttendanceHistoryService) DeleteServiceHistory(context context.Context, queries *db.Queries, deaconId int64, attendance payload.Attendance) error {
	dateOnly, err := time.Parse(time.DateOnly, attendance.Date)
	if err != nil {
		return err
	}
	return s.HistoryRepository.DeleteDeaconServiceHistory(context, queries, db.DeleteHistoryServiceParams{DeaconID: deaconId, Date: dateOnly, EventSkillLiturgyID: attendance.ESLId})
}

func (s AttendanceHistoryService) GetAllServiceHistory(context context.Context, queries *db.Queries, page int32, limit int32) ([]db.GetAllHistoryServiceRow, error) {
	return s.HistoryRepository.GetAllServiceHistory(context, queries, db.GetAllHistoryServiceParams{Offset: page, Limit: limit})
}

func (s AttendanceHistoryService) GetSuggestion(context context.Context, queries *db.Queries, eventSkillLiturgyId int32, page int32, limit int32) ([]db.GetSuggestionRow, error) {
	return s.HistoryRepository.GetSuggestion(context, queries, db.GetSuggestionParams{POffset: page, PLimit: limit, Eslid: eventSkillLiturgyId})
}
