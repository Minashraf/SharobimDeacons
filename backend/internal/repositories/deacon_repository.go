package repositories

import (
	db "Backend/internal/models"
	"context"
)

type DeaconRepository struct{}

func NewDeaconRepository() *DeaconRepository {
	return &DeaconRepository{}
}

func (deaconRepository *DeaconRepository) GetDeacons(context context.Context, queries *db.Queries, sorting map[string]string, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetDeaconByIdRow, error) {
	return queries.ListDeacons(context, sorting, deaconPage)
}

func (deaconRepository *DeaconRepository) GetDeaconProfile(context context.Context, queries *db.Queries, deaconId int64) (db.GetDeaconByIdRow, error) {
	return queries.GetDeaconById(context, deaconId)
}

func (deaconRepository *DeaconRepository) GetDeaconSkills(context context.Context, queries *db.Queries, deaconId int64) ([]db.GetDeaconSkillByIdRow, error) {
	return queries.GetDeaconSkillById(context, deaconId)
}

func (deaconRepository *DeaconRepository) GetDeaconServiceHistory(context context.Context, queries *db.Queries, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetHistoryServiceByDeaconIdRow, error) {
	return queries.GetHistoryServiceByDeaconId(context, deaconPage)
}

func (deaconRepository *DeaconRepository) AddDeacon(context context.Context, queries *db.Queries, deacon db.CreateDeaconParams) (int64, error) {
	return queries.CreateDeacon(context, deacon)
}

func (deaconRepository *DeaconRepository) AddDeaconSkill(context context.Context, queries *db.Queries, deaconSkill db.InsertDeaconSkillParams) error {
	return queries.InsertDeaconSkill(context, deaconSkill)
}

func (deaconRepository *DeaconRepository) DeleteDeacon(context context.Context, queries *db.Queries, deaconId int64) error {
	return queries.DeleteDeacon(context, deaconId)
}

func (deaconRepository *DeaconRepository) DeleteDeaconSkill(context context.Context, queries *db.Queries, deaconId int64) error {
	return queries.DeleteDeaconSkill(context, deaconId)
}

func (deaconRepository *DeaconRepository) UpdateDeacon(context context.Context, queries *db.Queries, params db.UpdateDeaconParams) error {
	return queries.UpdateDeacon(context, params)
}
