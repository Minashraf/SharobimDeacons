package routes

import (
	"Backend/internal/consts"
	"Backend/internal/handlers"
	"Backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User              *handlers.UserHandler
	Deacon            *handlers.DeaconHandler
	Skill             *handlers.SkillHandler
	Events            *handlers.EventsHandler
	Liturgy           *handlers.LiturgyHandler
	AttendanceHistory *handlers.AttendanceHistory
}

func (handler *Handler) Setup(router *gin.Engine) {
	handler.registerHandler()

	user := router.Group("/user")
	{
		user.POST("/register", handler.User.CreateUser)
		user.POST("/login", handler.User.Login)
		user.POST("/refresh", handler.User.Refresh)
	}

	deacons := router.Group("/deacons", middleware.AuthMiddleware(), middleware.AllowedRoles([]string{consts.SuperAdmin, consts.Admin}))
	{
		deacons.POST("/", handler.Deacon.CreateDeacon)
		deacons.PUT("/:id", handler.Deacon.UpdateDeacon)
		deacons.DELETE("/:id", handler.Deacon.DeleteDeacon)
		deacons.GET("/", handler.Deacon.GetDeacons)
		deacons.GET("/ranks", handler.Deacon.GetDeaconsRanks)
		deacons.GET("/:id", handler.Deacon.GetDeacon)
	}

	attendanceHistory := router.Group("/attendance", middleware.AuthMiddleware(), middleware.AllowedRoles([]string{consts.SuperAdmin, consts.Admin}))
	{
		deaconAttendance := attendanceHistory.Group("/deacon")
		{
			deaconAttendance.GET("/:id", handler.AttendanceHistory.GetServiceHistory)
			deaconAttendance.POST("/:id", handler.AttendanceHistory.AddServiceHistory)
			deaconAttendance.POST("/bulk", handler.AttendanceHistory.BulkAssign)
			deaconAttendance.DELETE("/:id", handler.AttendanceHistory.DeleteServiceHistory)
		}
		attendanceHistory.GET("/history", handler.AttendanceHistory.GetAllServiceHistory)
		attendanceHistory.GET("/suggestion", handler.AttendanceHistory.GetSuggestion)
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
	handler.AttendanceHistory = handlers.NewAttendanceHistoryHandler()
}
