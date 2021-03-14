package handlers

import (
	"net/http"
	"strconv"

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

var cls models.Client

// ShowClients test
func ShowClientsGorm(c echo.Context) error {
	response, err := cls.ShowClientGorm()
	if err != nil {
		panic(err)
	}
	// fmt.Println(response)
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"Title": "Clients",
		"clients": response,
	})
}

// SaveFormClient list client 
func SaveFormClient(c echo.Context) error {
	cli := new(models.Client)
	if err := c.Bind(cli); err != nil {
		return err
	}
	data := &models.Client{
		FirstName: cli.FirstName,
		LastName: cli.LastName,
		Ci: cli.Ci,
	}
	response, err := cls.CreateClientGorm(data)
	if err != nil {
		panic(err)
	}
	return c.Render(http.StatusOK, "msg.html", map[string]interface{}{
		"Title": "Result serach client",
		"msg": "Record saved",
		"client": response,
	})
}

// EditFormClient editclient
func EditFormClient (c echo.Context) error {
	tmp := c.Param("id")
	id, err := strconv.ParseInt(tmp, 10, 64)
	response, err := cls.EditClientGorm(id)
	if err != nil {
		panic(err)
	}
	return c.Render(http.StatusOK, "edit.html", map[string]interface{}{
		"Title": "Edit Client",
		"client": response,
	})
}

// UpdateClientGorm update client
func UpdateClientGorm(c echo.Context) error {
	cli := new(models.Client)
	if err := c.Bind(cli); err != nil {
		return err
	}

	data := &models.Client{
		FirstName: cli.FirstName,
		LastName: cli.LastName,
		Ci: cli.Ci,
	}

	tmp := c.Param("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		panic(err)
	}
	response, err := cls.SaveEditClientGorm(id, data)
	return c.Render(http.StatusOK, "msg.html", map[string]interface{}{
		"Title": "Record Saved",
		"client": response,
	})
}

// DeleteClientGorm delete a client
func DeleteClientGorm(c echo.Context) error {
	tmp := c.Param("id")
	id, err := strconv.Atoi(tmp)
	if err != nil {
		panic(err)
	}
	response := cls.DeleteClientGorm(id)
	return c.Render(http.StatusOK, "msg.html", map[string]interface{}{
		"Title": "Delete record",
		"client": response,
	})
}

// func EditFormClient (c echo.Context) error {
// 	tmp := c.Param("id")
// 	// pk, _ := strconv.Atoi(tmp)
// 	id, err := strconv.ParseInt(tmp, 10, 64)
// 	cls, err := models.EditClientGorm(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c.Render(http.StatusOK, "edit.html", map[string]interface{}{
// 		"Title": "Edit Client",
// 		"client": cls,
// 	})
// }