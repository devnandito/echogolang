package models

import (
	"fmt"
	"time"

	"github.com/devnandito/echogolang/lib"
)

// Client client access public
type Client struct {
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Ci string `json:"ci"`
	Birthday time.Time `json:"birthday"`
}

// BirthdayDateStr conver to string
func (c Client) BirthdayDateStr() string {
	return c.Birthday.Format("2006-01-02")
}
// BirthdayTime convert string to time
func (c Client) BirthdayTime(timeStr string) (timeT time.Time) {
	const Format = "2006-01-02"
	t, _ := time.Parse(Format, timeStr)
	return t
}

// SeekClient show all client
func SeekClient() ([]Client, error) {

	conn := lib.NewConfig()
	db := conn.DsnString()
	rows, err := db.Query("SELECT id, first_name, last_name, ci, birthday FROM clients LIMIT 10")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var cls []Client
	for rows.Next() {
		var cl Client
		err := rows.Scan(&cl.ID, &cl.FirstName, &cl.LastName, &cl.Ci, &cl.Birthday)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cl)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cls, nil
}

// CreateClient insert new client
func CreateClient(cls *Client)  (*Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnString()

	row, err := db.Prepare("INSERT INTO clients(first_name,last_name,ci,birthday) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err)
	}

	defer row.Close()

	row.Exec(cls.FirstName, cls.LastName, cls.Ci, cls.Birthday)
	fmt.Println(cls.FirstName, cls.LastName, cls.Ci, cls.Birthday)
	var i = &Client{
		FirstName: cls.FirstName,
		LastName: cls.LastName,
		Ci: cls.Ci,
		Birthday: cls.Birthday,
	}
	return i, nil
}

// UpdateClient update client
func UpdateClient(ci string, cls *Client)  (*Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnString()

	row, err := db.Prepare("UPDATE clients SET first_name = $1, last_name = $2 WHERE ci = $3")
	if err != nil {
		panic(err)
	}

	defer row.Close()

	row.Exec(cls.FirstName, cls.LastName, ci)
	fmt.Println(cls.FirstName, cls.LastName, ci)
	var i = &Client{
		FirstName: cls.FirstName,
		LastName: cls.LastName,
	}
	return i, nil
}

// DeleteClient delete client
func DeleteClient(ci string)  error {
	conn := lib.NewConfig()
	db := conn.DsnString()

	row, err := db.Prepare("DELETE FROM clients WHERE ci=$1")
	if err != nil {
		panic(err)
	}

	defer row.Close()
	row.Exec(ci)
	return nil
}

// GetClient search a client
func GetClient(ci int) ([]Client, error) {
	conn := lib.NewConfig()
	db := conn.DsnString()
	row, err := db.Query("SELECT id, first_name, last_name, ci, birthday FROM clients WHERE ci=$1", ci)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	var cls []Client
	for row.Next() {
		var cl Client
		err := row.Scan(&cl.ID, &cl.FirstName, &cl.LastName, &cl.Ci, &cl.Birthday)
		if err != nil {
			return nil, err
		}
		cls = append(cls, cl)
	}

	return cls, nil
}

// func CreateClient(cls *Client)  error {
// 	var LastInsertId int
// 	conn := lib.NewConfig()
// 	db := conn.DsnString()

// 	row := db.QueryRow("INSERT INTO clients(first_name,last_name,ci) VALUES($1,$2,$3) returning id;", cls.FirstName, cls.LastName, cls.LastName).Scan(&LastInsertId)
// 	fmt.Println(cls.FirstName, cls.LastName, cls.Ci)
// 	if row != nil {
// 		panic(row)
// 	}

// 	fmt.Println("Último id =", LastInsertId)
// 	return nil
// }