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
