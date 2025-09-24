package services

import (
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"context"
)

type LiturgyService struct {
	Repository *repositories.LiturgyRepository
}

func NewLiturgyService() *LiturgyService {
	return &LiturgyService{Repository: repositories.NewLiturgyRepository()}
}

func (s LiturgyService) GetLiturgies(context context.Context, queries *db.Queries) ([]db.GetLiturgiesRow, error) {
	return s.Repository.GetLiturgies(context, queries)
}
