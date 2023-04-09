package main

import (
	userHandler "github.com/dimasyudhana/latihan-deployment.git/app/features/users/handler"
	userRepo "github.com/dimasyudhana/latihan-deployment.git/app/features/users/repository"
	userLogic "github.com/dimasyudhana/latihan-deployment.git/app/features/users/service"
	"github.com/dimasyudhana/latihan-deployment.git/app/routes"
	"github.com/dimasyudhana/latihan-deployment.git/config"
	"github.com/labstack/echo/v4"
)

// PortNumber
const PortNumber = ":8080"

func main() {
	e := echo.New()
	// Database connection
	cfg := config.InitConfiguration()
	db, _ := config.GetConnection(*cfg)
	config.Migrate(db)

	userModel := userRepo.New(db)
	userService := userLogic.New(userModel)
	userController := userHandler.New(userService)

	// routing
	routes.Route(e, userController)

	if err := e.Start(PortNumber); err != nil {
		e.Logger.Fatal("cannot start server ", err.Error())
	}
}
