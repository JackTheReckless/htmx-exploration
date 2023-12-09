package forest

import (
	"htmx-exploration/enemy"
	"htmx-exploration/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ForestHandler struct {
	ForestService *ForestService
}

func NewForestHandler(forestService *ForestService) *ForestHandler {
	return &ForestHandler{
		ForestService: forestService,
	}
}

func (h *ForestHandler) Forest(c echo.Context) error {
	theUser := user.GetCurrentUser()

	newForest := NewForest()

	SetCurrentForest(newForest)

	data := map[string]interface{}{
		"User": theUser,
	}
	return c.Render(http.StatusOK, "forest", data)
}

func (h *ForestHandler) Search(c echo.Context) error {
	theForest := GetCurrentForest()

	theForest.RandomEncounter()

	encounterType := theForest.currentEncounter

	switch encounterType {
	case "basicCombat":
		return h.BasicCombat(c)
	default:
		return c.String(http.StatusBadRequest, "Unknown Encounter Type")
	}

}

func (h *ForestHandler) BasicCombat(c echo.Context) error {
	theUser := user.GetCurrentUser()
	theForest := GetCurrentForest()

	theForest.EncounterBasicEnemy()

	enemy := enemy.GetCurrentEnemy()

	data := map[string]interface{}{
		"User":  theUser,
		"Enemy": enemy,
	}

	return c.Render(http.StatusOK, "combat", data)
}
