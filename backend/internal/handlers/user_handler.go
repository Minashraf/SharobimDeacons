package handlers

import (
	"Backend/internal/data/payload"
	"Backend/internal/middleware"
	"Backend/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{services.NewUserService()}
}

func (userHandler *UserHandler) CreateUser(c *gin.Context) {
	var user payload.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := userHandler.Service.CreateUser(c.Request.Context(), middleware.GetQueries(c), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (userHandler *UserHandler) Login(c *gin.Context) {
	var user payload.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := userHandler.Service.Login(c.Request.Context(), middleware.GetQueries(c), &user)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	c.SetCookie(
		"jwt_token",
		token,
		int(24*time.Hour.Seconds()),
		"/",
		"localhost", //TODO actual domain
		true,
		true,
	)
	c.Status(http.StatusOK)
}
