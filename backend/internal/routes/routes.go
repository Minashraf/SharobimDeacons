package routes

import (
	"Backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User *handlers.UserHandler
}

func (handler *Handler) Setup(router *gin.Engine) {
	handler.registerHandler()

	user := router.Group("/user")
	{
		user.POST("/register", handler.User.CreateUser)
		user.POST("/login", handler.User.Login)
	}
}

func (handler *Handler) registerHandler() {
	handler.User = handlers.NewUserHandler()
}
