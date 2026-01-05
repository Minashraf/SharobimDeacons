package repositories

import (
	db "Backend/internal/models"
	"context"
)

type EventsRepository struct{}

func NewEventsRepository() *EventsRepository {
	return &EventsRepository{}
}

func (eventRepository *EventsRepository) GetEvents(context context.Context, queries *db.Queries, liturgyId int32) ([]db.GetEventsRow, error) {
	return queries.GetEvents(context, liturgyId)
}
