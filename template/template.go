package template

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	Templates map[string]*template.Template
}

// Render implement e.Renderer interface
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	tmpl, ok := t.Templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

// Templates ?
func Templates() map[string]*template.Template {
	t := make(map[string]*template.Template)
	t["home.html"] = template.Must(template.ParseFiles("views/home.html", "views/base.html"))
	t["about.html"] = template.Must(template.ParseFiles("views/about.html", "views/base.html"))
	return t
}
