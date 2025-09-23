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
