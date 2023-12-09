package combat

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, combatHandler *CombatHandler) {
	userGroup := e.Group("/combat")

	userGroup.GET("/actions", combatHandler.Actions)
	userGroup.GET("/log", combatHandler.CombatLog)
	userGroup.POST("/fight", combatHandler.Fight)
	userGroup.POST("/run", combatHandler.Run)
	userGroup.POST("/user/heal", combatHandler.UserHeal)
	userGroup.POST("/enemy/heal", combatHandler.EnemyHeal)
}
