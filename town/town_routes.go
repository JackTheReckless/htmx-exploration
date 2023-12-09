package town

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, townHandler *TownHandler) {
	townGroup := e.Group("/town")

	townGroup.GET("", townHandler.Town)
}
