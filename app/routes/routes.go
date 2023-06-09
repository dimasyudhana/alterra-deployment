package routes

import (
	"github.com/dimasyudhana/latihan-deployment.git/app/features/book"
	"github.com/dimasyudhana/latihan-deployment.git/app/features/user"
	"github.com/dimasyudhana/latihan-deployment.git/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, uc user.Handler, bc book.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())
	e.PUT("/users", uc.Update(), middleware.JWT([]byte(config.JWTSecret)))

	e.POST("/books", bc.InsertBook(), middleware.JWT([]byte(config.JWTSecret)))
	e.PUT("/books", bc.UpdateBook(), middleware.JWT([]byte(config.JWTSecret)))

}
