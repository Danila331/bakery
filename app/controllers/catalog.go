package controllers

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

func CatalogCakePageController(c echo.Context) error {
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

func CatalogCookiePageController(c echo.Context) error {
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
