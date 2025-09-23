package services

import (
	"Backend/internal/data/payload"
	"Backend/internal/data/response"
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"
)

type DeaconService struct {
	Repository *repositories.DeaconRepository
}

func NewDeaconService() *DeaconService {
	return &DeaconService{Repository: repositories.NewDeaconRepository()}
}

func (s DeaconService) GetDeacons(context context.Context, queries *db.Queries, sorting map[string]string, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetDeaconByIdRow, error) {
	return s.Repository.GetDeacons(context, queries, sorting, deaconPage)
}

func (s DeaconService) GetDeaconProfile(context context.Context, queries *db.Queries, deaconId int64) (response.GetDeaconById, error) {
	info, err := s.Repository.GetDeaconProfile(context, queries, deaconId)
	if err != nil {
		return response.GetDeaconById{}, err
	}
	skills, err := s.Repository.GetDeaconSkills(context, queries, deaconId)
	if err != nil {
		return response.GetDeaconById{}, err
	}
	return response.GetDeaconById{
		ID:          info.ID,
		FirstName:   info.FirstName,
		LastName:    info.LastName,
		DateOfBirth: info.DateOfBirth,
		Address:     info.Address,
		Email:       info.Email,
		PhoneNumber: info.PhoneNumber,
		Country:     info.Country,
		RankName:    info.RankName,
		Skills:      skills,
	}, nil
}

func (s DeaconService) GetServiceHistory(context context.Context, queries *db.Queries, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetHistoryServiceByDeaconIdRow, error) {
	return s.Repository.GetDeaconServiceHistory(context, queries, deaconPage)
}
func (s DeaconService) AddDeacon(context context.Context, queries *db.Queries, deacon payload.Deacon) error {
	var address sql.NullString
	if strings.TrimSpace(deacon.Address) == "" {
		address = sql.NullString{Valid: false}
	} else {
		address = sql.NullString{String: deacon.Address, Valid: true}
	}
	var email sql.NullString
	if strings.TrimSpace(deacon.Email) == "" {
		email = sql.NullString{Valid: false}
	} else {
		email = sql.NullString{String: deacon.Email, Valid: true}
	}
	var phone sql.NullString
	if strings.TrimSpace(deacon.Phone) == "" {
		phone = sql.NullString{Valid: false}
	} else {
		phone = sql.NullString{String: deacon.Phone, Valid: true}
	}
	var date sql.NullTime
	if strings.TrimSpace(deacon.DOB) == "" {
		date = sql.NullTime{Valid: false}
	} else {
		dateOnly, err := time.Parse(time.DateOnly, deacon.DOB)
		if err != nil {
			return err
		}
		date = sql.NullTime{Time: dateOnly, Valid: true}
	}

	params := db.CreateDeaconParams{
		FirstName:    deacon.FirstName,
		LastName:     deacon.LastName,
		Address:      address,
		Email:        email,
		PhoneNumber:  phone,
		DateOfBirth:  date,
		Country:      deacon.Country,
		DeaconRankID: deacon.DeaconRank,
	}
	deaconId, err := s.Repository.AddDeacon(context, queries, params)
	if err != nil {
		return err
	}
	for _, skill := range deacon.Skills {
		err = s.Repository.AddDeaconSkill(context, queries, db.InsertDeaconSkillParams{DeaconID: deaconId, SkillID: skill.Skill, Score: skill.Score})
		if err != nil {
			err2 := s.Repository.DeleteDeacon(context, queries, deaconId)
			if err2 != nil {
				return errors.Join(err, err2)
			}
			return err
		}
	}
	return nil
}

func (s DeaconService) DeleteDeacon(context context.Context, queries *db.Queries, deaconId int64) error {
	err := s.Repository.DeleteDeaconSkill(context, queries, deaconId)
	if err != nil {
		return err
	}
	err = s.Repository.DeleteDeacon(context, queries, deaconId)
	if err != nil {
		return err
	}
	return nil
}

func (s DeaconService) UpdateDeacon(context context.Context, queries *db.Queries, deacon payload.Deacon, deaconId int64) error {
	var address sql.NullString
	if strings.TrimSpace(deacon.Address) == "" {
		address = sql.NullString{Valid: false}
	} else {
		address = sql.NullString{String: deacon.Address, Valid: true}
	}
	var email sql.NullString
	if strings.TrimSpace(deacon.Email) == "" {
		email = sql.NullString{Valid: false}
	} else {
		email = sql.NullString{String: deacon.Email, Valid: true}
	}
	var phone sql.NullString
	if strings.TrimSpace(deacon.Phone) == "" {
		phone = sql.NullString{Valid: false}
	} else {
		phone = sql.NullString{String: deacon.Phone, Valid: true}
	}
	var date sql.NullTime
	if strings.TrimSpace(deacon.DOB) == "" {
		date = sql.NullTime{Valid: false}
	} else {
		dateOnly, err := time.Parse(time.DateOnly, deacon.DOB)
		if err != nil {
			return err
		}
		date = sql.NullTime{Time: dateOnly, Valid: true}
	}

	params := db.UpdateDeaconParams{
		FirstName:    deacon.FirstName,
		LastName:     deacon.LastName,
		Address:      address,
		Email:        email,
		PhoneNumber:  phone,
		DateOfBirth:  date,
		Country:      deacon.Country,
		DeaconRankID: deacon.DeaconRank,
		ID:           deaconId,
	}
	err := s.Repository.UpdateDeacon(context, queries, params)
	if err != nil {
		return err
	}
	err = s.Repository.DeleteDeaconSkill(context, queries, deaconId)
	if err != nil {
		return err
	}
	for _, skill := range deacon.Skills {
		err = s.Repository.AddDeaconSkill(context, queries, db.InsertDeaconSkillParams{DeaconID: deaconId, SkillID: skill.Skill, Score: skill.Score})
		if err != nil {
			return err
		}
	}
	return nil
}
