package repositories

import (
	db "Backend/internal/models"
	"context"
)

type LiturgyRepository struct{}

func NewLiturgyRepository() *LiturgyRepository {
	return &LiturgyRepository{}
}

func (liturgyRepository *LiturgyRepository) GetLiturgies(context context.Context, queries *db.Queries) ([]db.GetLiturgiesRow, error) {
	return queries.GetLiturgies(context)
}
