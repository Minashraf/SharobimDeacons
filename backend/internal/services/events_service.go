package services

import (
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"context"
)

type EventsService struct {
	Repository *repositories.EventsRepository
}

func NewEventsService() *EventsService {
	return &EventsService{Repository: repositories.NewEventsRepository()}
}

func (s EventsService) GetEvents(context context.Context, queries *db.Queries, liturgyId int32) ([]db.GetEventsRow, error) {
	return s.Repository.GetEvents(context, queries, liturgyId)
}
