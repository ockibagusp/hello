package template

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
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

func parseFiles(s string, t ...string) *template.Template {
	// t parseFiles, example "views/user-form.html"
	if len(t) == 1 {
		return template.Must(template.ParseFiles(s, t[0], "views/base.html"))
	} else if len(t) >= 2 {
		panic("t [1] parseFiles, example \"views/users/user-form.html\"")
	}
	// "views/base.html"?
	return template.Must(template.ParseFiles(s, "views/base.html"))
}

// Templates
func Templates() *TemplateRenderer {
	t := make(map[string]*template.Template)
	t["home.html"] = parseFiles("views/home.html")
	t["about.html"] = parseFiles("views/about.html")
	t["users/user-all.html"] = parseFiles("views/users/user-all.html")
	t["users/user-add.html"] = parseFiles("views/users/user-add.html", "views/users/user-form.html")
	t["users/user-read.html"] = parseFiles("views/users/user-read.html", "views/users/user-form.html")
	t["users/user-view.html"] = parseFiles("views/users/user-view.html", "views/users/user-form.html")

	return &TemplateRenderer{
		Templates: t,
	}
}
