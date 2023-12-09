package enemy

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EnemyHandler struct {
	EnemyService *EnemyService
}

func NewEnemyHandler(enemyService *EnemyService) *EnemyHandler {
	return &EnemyHandler{
		EnemyService: enemyService,
	}
}

func (h *EnemyHandler) Health(c echo.Context) error {
	theEnemy := GetCurrentEnemy()

	data := map[string]interface{}{
		"Enemy": theEnemy,
	}
	return c.Render(http.StatusOK, "combat-enemy", data)
}
