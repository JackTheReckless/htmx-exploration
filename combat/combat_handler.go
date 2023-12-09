package combat

import (
	"htmx-exploration/enemy"
	"htmx-exploration/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CombatHandler struct {
	CombatService *CombatService
}

func NewCombatHandler(combatService *CombatService) *CombatHandler {
	return &CombatHandler{
		CombatService: combatService,
	}
}

func (h *CombatHandler) Actions(c echo.Context) error {
	theUser := user.GetCurrentUser()

	data := map[string]interface{}{
		"User": theUser,
	}
	return c.Render(http.StatusOK, "combat-actions", data)
}

func (h *CombatHandler) CombatLog(c echo.Context) error {
	theCombatLog := GetCombatLog()

	data := map[string]interface{}{
		"Combat": theCombatLog,
	}

	return c.Render(http.StatusOK, "combat-log", data)
}

func (h *CombatHandler) Fight(c echo.Context) error {
	theUser := user.GetCurrentUser()
	theEnemy := enemy.GetCurrentEnemy()

	UserAttack(theUser, theEnemy)

	combatLog := GetCombatLog()

	data := map[string]interface{}{
		"User":   theUser,
		"Enemy":  theEnemy,
		"Combat": combatLog,
	}

	NewCombatRound()

	if theEnemy.Health == 0 {
		enemy.ResetEnemy()
		ResetCombat()
	}

	return c.Render(http.StatusOK, "combat", data)
}

func (h *CombatHandler) Run(c echo.Context) error {
	theUser := user.GetCurrentUser()
	ResetCombat()

	data := map[string]interface{}{
		"User": theUser,
	}
	return c.Render(http.StatusOK, "forest", data)
}

func (h *CombatHandler) UserHeal(c echo.Context) error {
	theUser := user.GetCurrentUser()

	UserHeal(theUser)
	NewCombatRound()

	c.Response().Header().Add("HX-Trigger", `{"user-change":"","combat-log":""}`)

	return c.String(http.StatusNoContent, "user healed")
}

func (h *CombatHandler) EnemyHeal(c echo.Context) error {
	theEnemy := enemy.GetCurrentEnemy()

	EnemyHeal(theEnemy)
	NewCombatRound()

	c.Response().Header().Add("HX-Trigger", `{"enemy-change":"","combat-log":""}`)

	return c.String(http.StatusNoContent, "enemy healed")
}
