package template

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Templates is a custom html/template renderer for Echo framework
type Templates struct {
	Templates map[string]*template.Template
}

// New Templates
func NewTemplates() *Templates {
	// TODO: tpl := template.Must(template.New("main").Funcs(template.FuncMap{"tostring": ToString}))

	t := make(map[string]*template.Template)
	t["home.html"] = parseFiles("views/home.html")
	t["about.html"] = parseFiles("views/about.html")
	// t["login.html"] = parseFiles("views/login.html") -> base.html
	t["login.html"] = template.Must(
		template.ParseFiles("views/login.html"),
	)
	t["users/user-all.html"] = parseFiles("views/users/user-all.html")
	t["users/user-add.html"] = parseFiles("views/users/user-add.html", "views/users/user-form.html")
	t["users/user-read.html"] = parseFiles("views/users/user-read.html", "views/users/user-form.html")
	t["users/user-view.html"] = parseFiles("views/users/user-view.html", "views/users/user-form.html")
	/*
		t["users/user-view-password.html"] -> no
			{..,"status":500,"error":"html/template: \"users/user-view-password.html\" is undefined",..}
			why?
		t["user-view-password.html"] -> yes
		----
		t["user-view-password.html"] = parseFiles("views/users/user-view-password.html") -> base.html
	*/
	t["user-view-password.html"] = template.Must(
		template.ParseFiles("views/users/user-view-password.html"),
	)

	return &Templates{
		Templates: t,
	}
}

// Render implement e.Renderer interface
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]
	if !ok {
		return errors.New("Template not found -> " + name)
	}

	// Add global methods if data is a map
	if viewContext, isMap := data.(echo.Map); isMap {
		viewContext["reverse"] = c.Echo().Reverse

		/*
			@param /login and /users/view/:id/password
			- Is HTML Only
			is_html_only (bool): {true, false}

			TODO:
			Login name: 				"login.html -> ok
			UpdateUserByPassword name: 	"user-view-password.html" -> ok
										"users/user-view-password.html" -> no
			why?
		*/
		if viewContext["is_html_only"] == true {
			return tmpl.ExecuteTemplate(w, name, data)
		}
	}

	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func parseFiles(s string, t ...string) *template.Template {
	// t parseFiles, example "views/user-form.html"
	if len(t) == 1 {
		return template.Must(
			/* template.New("") or template.New("base"), same ?

			t := make(map[string]*template.Template)
			t["home.html"] = parseFiles("views/home.html")
			...

			*/
			template.New("").Funcs(template.FuncMap{"tostring": ToString}).
				ParseFiles(s, t[0], "views/base.html"),
		)
	} else if len(t) >= 2 {
		panic("t [1] parseFiles, example \"views/users/user-form.html\"")
	}
	// "views/base.html"?
	return template.Must(
		template.New("").Funcs(template.FuncMap{"tostring": ToString}).
			ParseFiles(s, "views/base.html"),
	)
}
