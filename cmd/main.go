package main

import (
	"log"

	"github.com/tonoy30/go-commerce/internal/app"
	"github.com/tonoy30/go-commerce/internal/repository"
	"github.com/tonoy30/go-commerce/internal/service"
)

func main() {
	// DB connection
	db, err := repository.NewSQLDB()
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}
	// register all services
	dao := repository.NewSQLDao(db)
	courseService := service.NewCourseService(dao)

	// start server
	app := app.NewApp(courseService)
	app.Start()
}
