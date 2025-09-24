package routes

import (
	"Backend/internal/consts"
	"Backend/internal/handlers"
	"Backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User    *handlers.UserHandler
	Deacon  *handlers.DeaconHandler
	Skill   *handlers.SkillHandler
	Events  *handlers.EventsHandler
	Liturgy *handlers.LiturgyHandler
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

	skills := router.Group("/skills", middleware.AuthMiddleware(), middleware.AllowedRoles([]string{consts.SuperAdmin, consts.Admin}))
	{
		skills.GET("/", handler.Skill.GetSkills)
		skills.GET("/:liturgy_id/:event_id", handler.Skill.GetDependantSkills)
	}

	events := router.Group("/events", middleware.AuthMiddleware(), middleware.AllowedRoles([]string{consts.SuperAdmin, consts.Admin}))
	{
		events.GET("/:liturgy_id", handler.Events.GetEvents)
	}

	liturgy := router.Group("/liturgy", middleware.AuthMiddleware(), middleware.AllowedRoles([]string{consts.SuperAdmin, consts.Admin}))
	{
		liturgy.GET("/", handler.Liturgy.GetLiturgies)
	}

}

func (handler *Handler) registerHandler() {
	handler.User = handlers.NewUserHandler()
	handler.Deacon = handlers.NewDeaconHandler()
	handler.Skill = handlers.NewSkillHandler()
	handler.Events = handlers.NewEventsHandler()
	handler.Liturgy = handlers.NewLiturgyHandler()
}
