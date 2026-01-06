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

func (s UserService) CreateUser(context context.Context, queries *db.Queries, user *payload.User) (string, error) {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	params := db.CreateUserParams{
		Email:    html.EscapeString(strings.TrimSpace(user.Email)),
		Password: hashed,
	}
	userId, err := s.Repository.CreateUser(context, queries, params)
	if err != nil {
		return "", err
	}
	//TODO needs to be configured
	err = s.Repository.AssignRole(context, queries, db.AssignRoleParams{RoleID: 2, UserID: userId})
	if err != nil {
		return "", err
	}
	return s.Login(context, queries, user)
}

func (s UserService) Login(context context.Context, queries *db.Queries, user *payload.User) (string, error) {
	userEntity, err := s.Repository.GetUserByEmail(context, queries, user.Email)
	if err != nil {
		return "", err
	}
	if err = utils.CheckPasswordHash(user.Password, userEntity.Password); err != nil {
		return "", err
	}
	token, err := utils.CreateToken(userEntity)
	if err != nil {
		return "", err
	}
	return token, nil
}
