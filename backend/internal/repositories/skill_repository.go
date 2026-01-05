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

func (skillRepository *SkillRepository) GetDependantSkills(context context.Context, queries *db.Queries, params db.GetDependantSkillsParams) ([]db.GetDependantSkillsRow, error) {
	return queries.GetDependantSkills(context, params)
}
