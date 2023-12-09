package enemy

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, enemyHandler *EnemyHandler) {
	enemyGroup := e.Group("/enemy")

	enemyGroup.GET("/health", enemyHandler.Health)
}
