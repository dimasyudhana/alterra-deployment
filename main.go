package main

import (
	bookHandler "github.com/dimasyudhana/latihan-deployment.git/app/features/book/handlers"
	bookRepo "github.com/dimasyudhana/latihan-deployment.git/app/features/book/repository"
	bookLogic "github.com/dimasyudhana/latihan-deployment.git/app/features/book/usecase"
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

	bookModel := bookRepo.New(db)
	bookService := bookLogic.New(bookModel)
	bookController := bookHandler.New(bookService)

	routes.Route(e, userController, bookController)

	e.Start(":8080")
}
