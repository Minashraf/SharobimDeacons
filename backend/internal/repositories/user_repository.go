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

func (userRepository *UserRepository) GetUserByEmail(context context.Context, queries *db.Queries, email string) (db.GetUserByEmailRow, error) {
	return queries.GetUserByEmail(context, email)
}

func (userRepository *UserRepository) GetUserAndRolesById(context context.Context, queries *db.Queries, userID int64) (db.GetUserAndRolesByIdRow, error) {
	return queries.GetUserAndRolesById(context, userID)
}

func (userRepository *UserRepository) AddOrUpdateRefreshToken(context context.Context, queries *db.Queries, params db.AddOrUpdateRefreshTokenParams) error {
	return queries.AddOrUpdateRefreshToken(context, params)
}

func (userRepository *UserRepository) GetUserIdByRefreshToken(context context.Context, queries *db.Queries, refreshToken string) (int64, error) {
	return queries.GetUserIdByRefreshToken(context, refreshToken)
}
