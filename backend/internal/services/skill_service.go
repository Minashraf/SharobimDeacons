package services

import (
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"context"
)

type SkillService struct {
	Repository *repositories.SkillRepository
}

func NewSkillService() *SkillService {
	return &SkillService{Repository: repositories.NewSkillRepository()}
}

func (s SkillService) GetSkills(context context.Context, queries *db.Queries) ([]db.GetSkillsRow, error) {
	return s.Repository.GetSkills(context, queries)
}

func (s SkillService) GetDependantSkills(context context.Context, queries *db.Queries, params db.GetDependantSkillsParams) ([]db.GetDependantSkillsRow, error) {
	return s.Repository.GetDependantSkills(context, queries, params)
}
