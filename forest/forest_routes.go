package forest

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, forestHandler *ForestHandler) {
	forestGroup := e.Group("/forest")

	forestGroup.GET("", forestHandler.Forest)
	forestGroup.GET("/search", forestHandler.Search)
}
