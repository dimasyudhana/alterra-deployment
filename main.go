package main

import (
	"log"

	userHandler "github.com/dimasyudhana/latihan-deployment.git/app/features/users/handler"
	userRepo "github.com/dimasyudhana/latihan-deployment.git/app/features/users/repository"
	userLogic "github.com/dimasyudhana/latihan-deployment.git/app/features/users/service"
	"github.com/dimasyudhana/latihan-deployment.git/app/routes"
	"github.com/dimasyudhana/latihan-deployment.git/config"
	"github.com/labstack/echo/v4"
)

// const PortNumber = ":8080"

func main() {
	e := echo.New()
	// Database connection
	cfg := config.InitConfiguration()
	db, err := config.GetConnection(cfg)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	log.Println("Connected with database!")
	config.Migrate(db)

	userModel := userRepo.New(db)
	userService := userLogic.New(userModel)
	userController := userHandler.New(userService)

	// routing
	routes.Route(e, userController)

	e.Start(":8080")
}
