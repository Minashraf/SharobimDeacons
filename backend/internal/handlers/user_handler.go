package handlers

import (
	"Backend/internal/data/payload"
	"Backend/internal/data/response"
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
	token, err := userHandler.Service.CreateUser(c.Request.Context(), middleware.GetQueries(c), &user)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.LoginResponse{Token: token})
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
	token, err := userHandler.Service.Login(c.Request.Context(), middleware.GetQueries(c), &user)
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, response.LoginResponse{Token: token})
}
