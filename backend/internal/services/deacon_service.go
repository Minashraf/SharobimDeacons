package services

import (
	"Backend/internal/data/payload"
	"Backend/internal/data/response"
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"context"
	"database/sql"
	"strings"
	"time"
)

type DeaconService struct {
	DeaconRepository *repositories.DeaconRepository
	SkillRepository  *repositories.SkillRepository
}

func NewDeaconService() *DeaconService {
	return &DeaconService{DeaconRepository: repositories.NewDeaconRepository()}
}

func (s DeaconService) GetDeacons(context context.Context, queries *db.Queries, sorting map[string]string, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetDeaconByIdRow, error) {
	return s.DeaconRepository.GetDeacons(context, queries, sorting, deaconPage)
}

func (s DeaconService) GetDeaconsRanks(context context.Context, queries *db.Queries) ([]db.GetDeaconsRanksRow, error) {
	return s.DeaconRepository.GetDeaconsRanks(context, queries)
}

func (s DeaconService) GetDeaconProfile(context context.Context, queries *db.Queries, deaconId int64) (response.GetDeaconById, error) {
	info, err := s.DeaconRepository.GetDeaconProfile(context, queries, deaconId)
	if err != nil {
		return response.GetDeaconById{}, err
	}
	skills, err := s.DeaconRepository.GetDeaconSkills(context, queries, deaconId)
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

func (s DeaconService) AddDeacon(context context.Context, queries *db.Queries, database *sql.DB, deacon payload.Deacon) error {
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
		FirstName:    strings.TrimSpace(deacon.FirstName),
		LastName:     strings.TrimSpace(deacon.LastName),
		Address:      address,
		Email:        email,
		PhoneNumber:  phone,
		DateOfBirth:  date,
		Country:      deacon.Country,
		DeaconRankID: deacon.DeaconRank,
	}
	deaconSkillMap := make(map[int32]int32)
	for _, skill := range deacon.Skills {
		deaconSkillMap[skill.SkillID] = skill.Score
	}
	tx, err := database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := queries.WithTx(tx)

	deaconId, err := s.DeaconRepository.AddDeacon(context, qtx, params)
	if err != nil {
		return err
	}
	allSkills, err := s.SkillRepository.GetSkills(context, queries)
	if err != nil {
		return err
	}
	for _, skill := range allSkills {
		if score, exists := deaconSkillMap[skill.ID]; exists {
			err = s.DeaconRepository.AddDeaconSkill(context, qtx, db.InsertDeaconSkillParams{DeaconID: deaconId, SkillID: skill.ID, Score: score})
		} else {
			err = s.DeaconRepository.AddDeaconSkill(context, qtx, db.InsertDeaconSkillParams{DeaconID: deaconId, SkillID: skill.ID, Score: 0})
		}
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (s DeaconService) DeleteDeacon(context context.Context, queries *db.Queries, database *sql.DB, deaconId int64) error {
	tx, err := database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := queries.WithTx(tx)
	err = s.DeaconRepository.DeleteDeaconSkill(context, qtx, deaconId)
	if err != nil {
		return err
	}

	err = s.DeaconRepository.DeleteDeacon(context, qtx, deaconId)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s DeaconService) UpdateDeacon(context context.Context, queries *db.Queries, database *sql.DB, deacon payload.Deacon, deaconId int64) error {
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
	tx, err := database.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := queries.WithTx(tx)
	err = s.DeaconRepository.UpdateDeacon(context, qtx, params)
	if err != nil {
		return err
	}
	err = s.DeaconRepository.DeleteDeaconSkill(context, qtx, deaconId)
	if err != nil {
		return err
	}
	for _, skill := range deacon.Skills {
		err = s.DeaconRepository.AddDeaconSkill(context, qtx, db.InsertDeaconSkillParams{DeaconID: deaconId, SkillID: skill.SkillID, Score: skill.Score})
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
