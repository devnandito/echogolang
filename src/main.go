package main

import (
	"github.com/devnandito/echogolang/handlers"
	"github.com/labstack/echo"
)

func main() {
	// Instanciar echo
	e := echo.New()
	e.GET("/", handlers.Home)
	e.GET("/clients", handlers.GetAllClients)
	e.POST("/clients", handlers.CreateClient)
	e.PUT("/clients/:ci", handlers.UpdateClient)
	e.DELETE("/clients/:ci", handlers.DeleteClient)
	e.GET("/clients/:ci", handlers.SearchClient)
	e.Logger.Fatal(e.Start(":9000"))
}