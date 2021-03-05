package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/devnandito/echogolang/lib"
	"github.com/devnandito/echogolang/models"
	"github.com/labstack/echo"
)

func main() {
	// Instanciar echo
	a := lib.NewConfig()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		
		DB, err := sql.Open("postgres", a.ConnectionString())

		if err != nil {
			log.Fatal(err)
		}

		cls, err := models.AllClient()
		
		if err != nil {
			log.Println(err)
		}

		for _, cl := range cls {
			fmt.Println(cl.ID, cl.FirstName, cl.LastName, cl.Ci, cl.Birthday)
		}
		return c.String(http.StatusOK, "Hello World!")
		
	})
	e.Logger.Fatal(e.Start(":9000"))
}