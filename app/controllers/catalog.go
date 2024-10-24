package controllers

import (
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/Danila331/bakery/app/models"
	"github.com/labstack/echo/v4"
)

func CatalogCakesPageController(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "catalog-cakes.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "cakes", nil)
	return nil
}

func CatalogCookiesPageController(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "catalog-cookies.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "cookies", nil)
	return nil
}

func CatalogCakePageController(c echo.Context) error {
	idParam := c.QueryParam("id")
	idParam_int, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	cake := models.Cakes[idParam_int-1]
	htmlFiles := []string{
		filepath.Join("./", "templates", "card.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "card", cake)
	return nil
}

func CatalogCookiePageController(c echo.Context) error {
	idParam := c.QueryParam("id")
	idParam_int, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	cake := models.Cookies[idParam_int-1]
	htmlFiles := []string{
		filepath.Join("./", "templates", "card.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "card", cake)
	return nil
}
