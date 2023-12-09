package main

import (
	"fmt"
	"html/template"
	"htmx-exploration/combat"
	"htmx-exploration/enemy"
	"htmx-exploration/forest"
	"htmx-exploration/town"
	"htmx-exploration/user"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template Renderer is a custom renderer for HTML templates
type TemplateRenderer struct {
	templates *template.Template
}

// Render implements the echo.Renderer interface
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files from the "static" directory
	e.Static("/static", "static")

	// Load HTML templates from user and town packages
	templates := template.Must(template.ParseGlob("templates/*.html"))
	templates = template.Must(templates.ParseGlob("town/templates/*html"))
	templates = template.Must(templates.ParseGlob("user/templates/*html"))
	templates = template.Must(templates.ParseGlob("combat/templates/*html"))
	templates = template.Must(templates.ParseGlob("forest/templates/*html"))

	// Register the HTML template renderer
	renderer := &TemplateRenderer{
		templates: templates,
	}

	e.Renderer = renderer

	// Print the names of loaded templates for debugging
	for _, t := range templates.Templates() {
		fmt.Println("Template:", t.Name())
	}

	// Define a route to render the main page
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	townService := town.NewTownService()
	userService := user.NewUserService()
	enemyService := enemy.NewEnemyService()
	combatService := combat.NewCombatService()
	forestService := forest.NewForestService()

	townHandler := town.NewTownHandler(townService)
	userHandler := user.NewUserHandler(userService)
	enemyHandler := enemy.NewEnemyHandler(enemyService)
	combatHandler := combat.NewCombatHandler(combatService)
	forestHandler := forest.NewForestHandler(forestService)

	town.RegisterRoutes(e, townHandler)
	user.RegisterRoutes(e, userHandler)
	enemy.RegisterRoutes(e, enemyHandler)
	combat.RegisterRoutes(e, combatHandler)
	forest.RegisterRoutes(e, forestHandler)

	e.POST("/sign-out", func(c echo.Context) error {
		user.DeleteUser()
		forest.ResetForest()

		return c.Render(http.StatusOK, "index.html", nil)
	})

	// Start the server with Air
	e.Start(":8080")
}
