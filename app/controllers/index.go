package controllers

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

func MainPageController(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "index.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "index", nil)
	return nil
}
