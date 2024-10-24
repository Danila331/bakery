package servers

import (
	"github.com/Danila331/bakery/app/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "./static")

	e.GET("/", controllers.MainPageController)
	catalog := e.Group("/catalog")
	catalog.GET("/cakes", controllers.CatalogCakesPageController)
	catalog.GET("/cookies", controllers.CatalogCookiesPageController)
	catalog.GET("/cakes/", controllers.CatalogCakePageController)
	catalog.GET("/cookies/", controllers.CatalogCookiePageController)
	e.Logger.Fatal(e.Start(":8080"))
	return nil
}
