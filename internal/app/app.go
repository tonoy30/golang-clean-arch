package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tonoy30/clean-arch/internal/service"
	"log"
)

type App interface {
	Serve()
}

type app struct {
	server        *echo.Echo
	courseService service.CourseService
}

func NewApp(courseService service.CourseService) App {
	server := echo.New()
	server.Use(middleware.CORS())
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())

	return &app{
		server:        server,
		courseService: courseService,
	}
}

func (a app) ConfigureRoutes() {
	a.server.GET("/courses/:id", a.GetCourseById)
	a.server.GET("/healthy", a.HealthCheck)
}
func (a app) Serve() {
	a.ConfigureRoutes()
	log.Println("Listening on port 5050")
	err := a.server.Start(":5050")
	if err != nil {
		log.Fatal("Something wrong with serving api")
	}
}
