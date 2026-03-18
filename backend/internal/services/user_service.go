package services

import (
	"Backend/internal/data/payload"
	"Backend/internal/data/response"
	db "Backend/internal/models"
	"Backend/internal/repositories"
	"Backend/internal/utils"
	"context"
	"golang.org/x/net/html"
	"strings"
	"time"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{Repository: repositories.NewUserRepository()}
}

func (s UserService) CreateUser(context context.Context, queries *db.Queries, user *payload.User) (response.LoginResponse, error) {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return response.LoginResponse{}, err
	}
	params := db.CreateUserParams{
		Email:    html.EscapeString(strings.TrimSpace(user.Email)),
		Password: hashed,
	}
	userId, err := s.Repository.CreateUser(context, queries, params)
	if err != nil {
		return response.LoginResponse{}, err
	}
	//TODO needs to be configured
	err = s.Repository.AssignRole(context, queries, db.AssignRoleParams{RoleID: 2, UserID: userId})
	if err != nil {
		return response.LoginResponse{}, err
	}
	return s.Login(context, queries, user)
}

func (s UserService) Login(context context.Context, queries *db.Queries, user *payload.User) (response.LoginResponse, error) {
	userEntity, err := s.Repository.GetUserByEmail(context, queries, user.Email)
	if err != nil {
		return response.LoginResponse{}, err
	}
	if err = utils.CheckPasswordHash(user.Password, userEntity.Password); err != nil {
		return response.LoginResponse{}, err
	}
	token, err := utils.CreateToken(userEntity)
	if err != nil {
		return response.LoginResponse{}, err
	}
	refreshToken, err := utils.GenerateRefreshToken()
	err = s.Repository.AddOrUpdateRefreshToken(context, queries, db.AddOrUpdateRefreshTokenParams{UserID: userEntity.ID, RefreshToken: utils.HashToken(refreshToken), UpdatedAt: time.Now()})
	if err != nil {
		return response.LoginResponse{}, err
	}
	return response.LoginResponse{Token: token, RefreshToken: refreshToken}, nil
}

func (s UserService) RefreshToken(context context.Context, queries *db.Queries, request *payload.RefreshPayload) (response.LoginResponse, error) {
	userID, err := s.Repository.GetUserIdByRefreshToken(context, queries, utils.HashToken(request.RefreshToken))
	if err != nil {
		return response.LoginResponse{}, err
	}
	userEntity, err := s.Repository.GetUserAndRolesById(context, queries, userID)
	if err != nil {
		return response.LoginResponse{}, err
	}
	token, err := utils.CreateToken(db.GetUserByEmailRow{ID: userEntity.ID, Roles: userEntity.Roles})
	if err != nil {
		return response.LoginResponse{}, err
	}
	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return response.LoginResponse{}, err
	}
	err = s.Repository.AddOrUpdateRefreshToken(context, queries, db.AddOrUpdateRefreshTokenParams{UserID: userEntity.ID, RefreshToken: utils.HashToken(refreshToken), UpdatedAt: time.Now()})
	if err != nil {
		return response.LoginResponse{}, err
	}
	return response.LoginResponse{Token: token, RefreshToken: refreshToken}, nil
}
