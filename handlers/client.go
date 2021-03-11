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