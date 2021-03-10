package main

import (
	"errors"
	"io"
	"text/template"

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
	e.GET("/", handlers.Home)
	e.GET("/clients", handlers.GetAllClients)
	e.GET("/clients/list", handlers.GetAllClientsGorm)
	e.POST("/clients", handlers.CreateClient)
	e.PUT("/clients/:ci", handlers.UpdateClient)
	e.DELETE("/clients/:ci", handlers.DeleteClient)
	e.GET("/clients/:ci", handlers.SearchClient)

	templates := make(map[string]*template.Template)
	templates["index.html"] = template.Must(template.ParseFiles("views/clients/index.html", "views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}
	e.GET("/clients/show", handlers.ShowClients)
	e.Static("/clients/static", "assets")
	e.Logger.Fatal(e.Start(":9000"))
}