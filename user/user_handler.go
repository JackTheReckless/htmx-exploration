package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) CreatingUser(c echo.Context) error {
	return c.Render(http.StatusOK, "create-user", nil)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	class := c.FormValue("class")

	newUser := NewUser(name, class)

	SetCurrentUser(newUser)
	theUser := GetCurrentUser()

	data := map[string]interface{}{
		"User": theUser,
	}

	return c.Render(http.StatusOK, "town", data)
}

func (h *UserHandler) ChangeClass(c echo.Context) error {
	class := c.FormValue("class")

	theUser := GetCurrentUser()

	theUser.SetClass(class)

	c.Response().Header().Add("HX-Trigger", `{"class-change":"","user-change":""}`)

	return c.String(http.StatusNoContent, "class changed successfully")
}

func (h *UserHandler) ChangeStat(c echo.Context) error {
	healthValue := c.FormValue("Health")
	staminaValue := c.FormValue("Stamina")
	theUser := GetCurrentUser()

	if healthValue != "" {
		health, err := strconv.Atoi(healthValue)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid value for Health")
		}

		theUser.SetCurrentHealth(health)
	}

	if staminaValue != "" {
		stamina, err := strconv.Atoi(staminaValue)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid value for Health")
		}

		theUser.SetCurrentStamina(stamina)
	}

	c.Response().Header().Add("HX-Trigger", "user-change")

	return c.String(http.StatusNoContent, "stat changed successfully")
}

func (h *UserHandler) GetUserInfo(c echo.Context) error {
	theUser := GetCurrentUser()

	data := map[string]interface{}{
		"User": theUser,
	}
	return c.Render(http.StatusOK, "user-info", data)
}

func (h *UserHandler) GetHealth(c echo.Context) error {
	theUser := GetCurrentUser()

	data := map[string]interface{}{
		"User": theUser,
	}
	return c.Render(http.StatusOK, "combat-user", data)
}

func (h *UserHandler) GetUserUpdate(c echo.Context) error {

	c.Response().Header().Add("HX-Trigger", "user-change")
	return c.String(http.StatusNoContent, "sent user-change event")
}
