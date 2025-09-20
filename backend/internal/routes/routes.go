package routes

import (
	"Backend/internal/consts"
	"Backend/internal/handlers"
	"Backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User   *handlers.UserHandler
	Deacon *handlers.DeaconHandler
}

func (handler *Handler) Setup(router *gin.Engine) {
	handler.registerHandler()

	user := router.Group("/user")
	{
		user.POST("/register", handler.User.CreateUser)
		user.POST("/login", handler.User.Login)
	}

	deacons := router.Group("/deacons", middleware.AuthMiddleware(), middleware.AllowedRoles([]string{consts.SuperAdmin}))
	{
		deacons.GET("/", handler.Deacon.GetDeacons)
		deacons.GET("/:id", handler.Deacon.GetProfile)
		deacons.GET("/history/:id", handler.Deacon.GetServiceHistory)
	}

}

func (handler *Handler) registerHandler() {
	handler.User = handlers.NewUserHandler()
	handler.Deacon = handlers.NewDeaconHandler()
}
