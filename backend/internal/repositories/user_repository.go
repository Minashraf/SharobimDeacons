package repositories

import (
	db "Backend/internal/models"
	"context"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (userRepository *UserRepository) CreateUser(context context.Context, queries *db.Queries, params db.CreateUserParams) (int64, error) {
	return queries.CreateUser(context, params)
}

func (userRepository *UserRepository) AssignRole(context context.Context, queries *db.Queries, params db.AssignRoleParams) error {
	return queries.AssignRole(context, params)
}
