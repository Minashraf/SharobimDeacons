package repositories

import (
	db "Backend/internal/models"
	"context"
)

type DeaconRepository struct{}

func NewDeaconRepository() *DeaconRepository {
	return &DeaconRepository{}
}

func (deaconRepository *DeaconRepository) GetDeaconProfile(context context.Context, queries *db.Queries, deaconId int64) (db.GetDeaconByIdRow, error) {
	return queries.GetDeaconById(context, deaconId)
}

func (deaconRepository *DeaconRepository) GetDeaconServiceHistory(context context.Context, queries *db.Queries, deaconPage db.GetHistoryServiceByDeaconIdParams) ([]db.GetHistoryServiceByDeaconIdRow, error) {
	return queries.GetHistoryServiceByDeaconId(context, deaconPage)
}
