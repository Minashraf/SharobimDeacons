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

	deacons := router.Group("/deacons", middleware.AuthMiddleware(), middleware.AllowedRoles([]string{consts.SuperAdmin, consts.Admin}))
	{
		deacons.POST("/", handler.Deacon.CreateDeacon)
		deacons.PUT("/:id", handler.Deacon.UpdateDeacon)
		deacons.DELETE("/:id", handler.Deacon.DeleteDeacon)
		deacons.GET("/", handler.Deacon.GetDeacons)
		deacons.GET("/:id", handler.Deacon.GetDeacon)
		deacons.GET("/:id/history", handler.Deacon.GetServiceHistory)
	}

}

func (handler *Handler) registerHandler() {
	handler.User = handlers.NewUserHandler()
	handler.Deacon = handlers.NewDeaconHandler()
}
