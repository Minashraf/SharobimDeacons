package services

import (
	"Backend/internal/data/payload"
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"Backend/internal/utils"
	"context"
	"golang.org/x/net/html"
	"strings"
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
		Email:    html.EscapeString(strings.TrimSpace(user.Email)),
		Password: hashed,
	}
	userId, err := s.Repository.CreateUser(context, queries, params)
	if err != nil {
		return err
	}
	//TODO needs to be configured
	return s.Repository.AssignRole(context, queries, db.AssignRoleParams{RoleID: 1, UserID: userId})
}

func (s UserService) Login(context context.Context, queries *db.Queries, user *payload.User) (string, error) {
	userEntity, err := s.Repository.GetUserByEmail(context, queries, user.Email)
	if err != nil {
		return "", err
	}
	if !utils.CheckPasswordHash(user.Password, userEntity.Password) {
		return "", err
	}
	token, err := utils.CreateToken(userEntity)
	if err != nil {
		return "", err
	}
	return token, nil
}
