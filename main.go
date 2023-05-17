package main

import (
	"html/template"
	"io"

	"syrlight/go-blog/config"
	"syrlight/go-blog/controllers"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func init() {
	config.GetDatabaseConnection()
}

func main() {

	// template setup
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	defer config.DB.Close()

	// New Echo Instance
	e := echo.New()
	e.Renderer = t

	// Routes
	e.GET("/", controllers.PostIndex)

	// Boostrap
	e.Logger.Fatal(e.Start(":3000"))
}
