package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/tonoy30/go-commerce/internal/service"
)

type App struct {
	server        *echo.Echo
	courseService service.CourseService
}

func NewApp(courseService service.CourseService) *App {
	server := echo.New()
	return &App{
		server:        server,
		courseService: courseService,
	}
}

func (a App) ConfigureRoutes() {
	a.server.GET("/", func(c echo.Context) error {
		userId := "TONOY"
		return c.String(200, userId)
	})
	a.server.GET("/healthy", a.HealthCheck)
}
func (a App) Start() {
	a.ConfigureRoutes()
	log.Println("Listening on port 5050")
	err := a.server.Start(":5050")
	if err != nil {
		log.Fatal("Something wrong with serving api")
	}
}
