package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/OckiFals/hello/handler"
	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates map[string]*template.Template
}

// Render implement e.Renderer interface
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)

}

func main() {
	// Echo instance
	e := echo.New()

	// Instantiate a template registry with an array of template set
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("views/home.html", "views/base.html"))
	templates["about.html"] = template.Must(template.ParseFiles("views/about.html", "views/base.html"))
	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	// // Why bootstrap.min.css, bootstrap.min.js, jquery.min.js?
	// http.Handle("/", http.FileServer(http.Dir("./assets/css")))
	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Route => handler
	e.GET("/", handler.HomeHandler).Name = "home"
	e.GET("/about", handler.AboutHandler).Name = "about"

	// Start the Echo server
	e.Start(":8000")
}
