package user

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, userHandler *UserHandler) {
	userGroup := e.Group("/user")

	userGroup.GET("/create", userHandler.CreatingUser)
	userGroup.GET("/info", userHandler.GetUserInfo)
	userGroup.GET("/health", userHandler.GetHealth)
	userGroup.GET("/stamina", userHandler.GetUserUpdate)
	userGroup.POST("/create", userHandler.CreateUser)
	userGroup.POST("/class-change", userHandler.ChangeClass)
	userGroup.POST("/stat-change", userHandler.ChangeStat)
}
