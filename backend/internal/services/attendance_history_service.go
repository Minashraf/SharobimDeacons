package services

import (
	"Backend/internal/data/payload"
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"context"
	"errors"
	"time"
)

type AttendanceHistoryService struct {
	HistoryRepository *repositories.AttendanceHistoryRepository
	DeaconRepository  *repositories.DeaconRepository
	ESLRepository     *repositories.EventSkillLiturgyRepository
}

func NewAttendanceHistoryServiceService() *AttendanceHistoryService {
	return &AttendanceHistoryService{HistoryRepository: repositories.NewAttendanceHistoryRepository(), DeaconRepository: repositories.NewDeaconRepository(), ESLRepository: repositories.NewEventSkillLiturgyRepository()}
}

func (s AttendanceHistoryService) GetServiceHistory(context context.Context, queries *db.Queries, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetHistoryServiceByDeaconIdRow, error) {
	return s.HistoryRepository.GetDeaconServiceHistory(context, queries, deaconPage)
}

func (s AttendanceHistoryService) AddServiceHistory(context context.Context, queries *db.Queries, deaconId int64, attendance payload.Attendance) error {
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
			return s.HistoryRepository.AddDeaconServiceHistory(context, queries, db.AddHistoryServiceParams{DeaconID: deaconId, Date: dateOnly, EventSkillLiturgyID: attendance.ESLId})
		}
	}
	return errors.New("the Deacon doesn't have the required Skill")
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
