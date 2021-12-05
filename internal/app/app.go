package app

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tonoy30/clean-arch/internal/service"
	"io/ioutil"
	"log"
)

type App interface {
	Serve()
}

type app struct {
	server        *echo.Echo
	authService   service.AuthService
	courseService service.CourseService
}

func NewApp(courseService service.CourseService, authService service.AuthService) App {
	server := echo.New()
	server.Use(middleware.CORS())
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())

	return &app{
		server:        server,
		authService:   authService,
		courseService: courseService,
	}
}

func (a app) ConfigureRoutes() {
	a.server.GET("/healthy", a.HealthCheck)

	auth := a.server.Group("/identity")
	auth.POST("/signup", a.SignUp)

	courses := a.server.Group("/courses")
	courses.GET("/:id", a.GetCourseById)

}
func (a app) Serve() {
	a.ConfigureRoutes()
	data, _ := json.MarshalIndent(a.server.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)
	log.Println("Listening on port 5050")
	err := a.server.Start(":5050")
	if err != nil {
		log.Fatal("Something wrong with serving api")
	}
}
