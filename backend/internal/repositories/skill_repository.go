package repositories

import (
	db "Backend/internal/models"
	"context"
)

type SkillRepository struct{}

func NewSkillRepository() *SkillRepository {
	return &SkillRepository{}
}

func (skillRepository *SkillRepository) GetSkills(context context.Context, queries *db.Queries) ([]db.GetSkillsRow, error) {
	return queries.GetSkills(context)
}
