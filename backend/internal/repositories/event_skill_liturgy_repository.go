package repositories

import (
	db "Backend/internal/models"
	"context"
)

type EventSkillLiturgyRepository struct{}

func NewEventSkillLiturgyRepository() *EventSkillLiturgyRepository {
	return &EventSkillLiturgyRepository{}
}

func (EventSkillLiturgy *EventSkillLiturgyRepository) GetEventSkillLiturgy(context context.Context, queries *db.Queries, eventSkillLiturgyId int32) (db.GetEventSkillLiturgyRow, error) {
	return queries.GetEventSkillLiturgy(context, eventSkillLiturgyId)
}
