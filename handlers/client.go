package handlers

import (
	"net/http"

	"github.com/devnandito/echogolang/models"
	"github.com/labstack/echo"
)

// ShowClients test
func ShowClients(c echo.Context) error {
	cls, err := models.SeekClient()
	if err != nil {
		panic(err)
	}
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Title": "Clients",
		"clients": cls,
	})
}

// SearchForm render form
func SearchForm(c echo.Context) error {
	return c.Render(http.StatusOK, "search.html", map[string]interface{}{
		"Title": "Search Form",
	})
}

// ResultSearch lista client 
func ResultSearch(c echo.Context) error {
	document :=c.FormValue("document")
	firstname := c.FormValue("first_name")
	lastname := c.FormValue("last_name")
	cls, err := models.GetClientGorm(document, firstname, lastname)
	if err != nil {
		panic(err)
	}
	
	return c.Render(http.StatusOK, "result.html", map[string]interface{}{
		"Title": "Result serach client",
		"clients": cls,
	})
}

// ShowFormClient render client form
func ShowFormClient(c echo.Context) error {
	return c.Render(http.StatusOK, "create.html", map[string]interface{}{
		"Title": "Create Form Client",
	})
}

// SaveFormClient lista client 
func SaveFormClient(c echo.Context) error {
	document :=c.FormValue("document")
	firstname := c.FormValue("first_name")
	lastname := c.FormValue("last_name")
	cls, err := models.CreateClientGorm(document, firstname, lastname)
	if err != nil {
		panic(err)
	}
	return c.Render(http.StatusOK, "msg.html", map[string]interface{}{
		"Title": "Result serach client",
		"msg": "Record saved",
		"client": cls,
	})
}