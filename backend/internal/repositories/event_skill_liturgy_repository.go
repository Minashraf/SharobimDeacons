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

func (EventSkillLiturgy *EventSkillLiturgyRepository) GetRemainingEventSkillLiturgyCapacity(context context.Context, queries *db.Queries, params db.GetRemainingEventSkillLiturgyCapacityParams) (int32, error) {
	return queries.GetRemainingEventSkillLiturgyCapacity(context, params)
}
