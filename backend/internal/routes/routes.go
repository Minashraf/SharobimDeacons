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
		user.POST("/", handler.User.CreateUser)
	}
}

func (handler *Handler) registerHandler() {
	handler.User = handlers.NewUserHandler()
}
