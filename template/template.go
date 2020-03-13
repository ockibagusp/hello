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

func parseFiles(s string) *template.Template {
	// "views/base.html"?
	return template.Must(template.ParseFiles(s, "views/base.html"))
}

// Templates ?
func Templates() *TemplateRenderer {
	t := make(map[string]*template.Template)
	t["home.html"] = parseFiles("views/home.html")
	t["about.html"] = parseFiles("views/about.html")
	t["user-all.html"] = parseFiles("views/user-all.html")
	t["user-add.html"] = parseFiles("views/user-add.html")
	t["user-read.html"] = parseFiles("views/user-read.html")
	t["user-view.html"] = parseFiles("views/user-view.html")

	return &TemplateRenderer{
		Templates: t,
	}
}
