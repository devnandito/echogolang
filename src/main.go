package main

import (
	"errors"
	"io"
	"text/template"

	"github.com/devnandito/echogolang/api"
	"github.com/devnandito/echogolang/handlers"
	"github.com/labstack/echo"
)

// TemplateRegistry initial
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Render template
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found"+ name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	// Instanciar echo
	e := echo.New()
	e.GET("/clients", api.GetAllClients)
	e.GET("/clients/list", api.GetAllClientsGorm)
	e.POST("/clients", api.CreateClient)
	e.PUT("/clients/:ci", api.UpdateClient)
	e.DELETE("/clients/:ci", api.DeleteClient)
	e.GET("/clients/:ci", api.SearchClient)
	
	templates := make(map[string]*template.Template)
	templates["index.html"] = template.Must(template.ParseFiles("views/clients/index.html", "views/base.html"))
	templates["home.html"] = template.Must(template.ParseFiles("views/home/index.html", "views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}
	e.Static("/static", "assets")
	e.File("/favicon.png", "static/img/favicon.png")
	e.GET("/", handlers.Home)
	e.Static("/clients/static", "assets")
	e.GET("/clients/show", handlers.ShowClients)
	e.Logger.Fatal(e.Start(":9000"))
}