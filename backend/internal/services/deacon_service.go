package services

import (
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"context"
)

type DeaconService struct {
	Repository *repositories.DeaconRepository
}

func NewDeaconService() *DeaconService {
	return &DeaconService{Repository: repositories.NewDeaconRepository()}
}

func (s DeaconService) GetDeaconProfile(context context.Context, queries *db.Queries, deaconId int64) (db.GetDeaconByIdRow, error) {
	return s.Repository.GetDeaconProfile(context, queries, deaconId)
}
