package main

import (
	userHandler "github.com/dimasyudhana/latihan-deployment.git/app/features/user/handlers"
	userRepo "github.com/dimasyudhana/latihan-deployment.git/app/features/user/repository"
	userLogic "github.com/dimasyudhana/latihan-deployment.git/app/features/user/usecase"
	"github.com/dimasyudhana/latihan-deployment.git/app/routes"
	"github.com/dimasyudhana/latihan-deployment.git/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfiguration()
	db, _ := config.GetConnection(cfg)
	config.Migrate(db)

	userModel := userRepo.New(db)
	userService := userLogic.New(userModel)
	userController := userHandler.New(userService)

	routes.Route(e, userController)

	e.Start(":8080")
}
