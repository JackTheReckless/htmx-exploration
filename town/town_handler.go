package town

import (
	"htmx-exploration/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TownHandler struct {
	TownService *TownService
}

func NewTownHandler(townService *TownService) *TownHandler {
	return &TownHandler{
		TownService: townService,
	}
}

func (h *TownHandler) Town(c echo.Context) error {
	theUser := user.GetCurrentUser()

	data := map[string]interface{}{
		"User": theUser,
	}
	return c.Render(http.StatusOK, "town", data)
}
