package handlers

import (
	"Backend/internal/data/payload"
	"Backend/internal/middleware"
	"Backend/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{services.NewUserService()}
}

// CreateUser @Summary User Creation
// @Description Creates a User for the application
// @Tags User
// @Accept json
// @Param payload.User body payload.User true "Creation"
// @Success 201 {object} response.LoginResponse
// @Failure 500
// @Failure 400
// @Router /user/register [post]
func (userHandler *UserHandler) CreateUser(c *gin.Context) {
	var user payload.User
	if err := c.BindJSON(&user); err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginResponse, err := userHandler.Service.CreateUser(c.Request.Context(), middleware.GetQueries(c), &user)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, loginResponse)
}

// Login @Summary User Login
// @Description Logs In The User
// @Tags User
// @Accept json
// @Produce json
// @Param payload.User body payload.User true "Logging"
// @Success 200 {object} response.LoginResponse
// @Failure 401
// @Failure 400
// @Router /user/login [post]
func (userHandler *UserHandler) Login(c *gin.Context) {
	var user payload.User
	if err := c.BindJSON(&user); err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginResponse, err := userHandler.Service.Login(c.Request.Context(), middleware.GetQueries(c), &user)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, loginResponse)
}

// Refresh @Summary User Refresh token
// @Description send refresh token to get new JWT Token
// @Tags User
// @Accept json
// @Produce json
// @Param payload.RefreshPayload body payload.RefreshPayload true "Refreshing"
// @Success 200 {object} response.LoginResponse
// @Failure 401
// @Failure 400
// @Router /user/refresh [post]
func (userHandler *UserHandler) Refresh(c *gin.Context) {
	var refreshPayload payload.RefreshPayload
	if err := c.BindJSON(&refreshPayload); err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginResponse, err := userHandler.Service.RefreshToken(c.Request.Context(), middleware.GetQueries(c), &refreshPayload)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, loginResponse)
}
