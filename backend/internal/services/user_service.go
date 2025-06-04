package services

import (
	"Backend/internal/data/payload"
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"Backend/internal/utils"
	"context"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{Repository: repositories.NewUserRepository()}
}

func (s UserService) CreateUser(context context.Context, queries *db.Queries, user *payload.User) error {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	params := db.CreateUserParams{
		Email:    user.Email,
		Password: hashed,
	}
	userId, err := s.Repository.CreateUser(context, queries, params)
	if err != nil {
		return err
	}
	return s.Repository.AssignRole(context, queries, db.AssignRoleParams{RoleID: 4, UserID: userId})
}
