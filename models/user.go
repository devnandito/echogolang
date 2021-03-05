package models

import (
	"database/sql"
	"time"
)

// DB connection
var DB *sql.DB

// Client client access public
type Client struct {
	ID int
	FirstName string
	LastName string
	Ci string
	Birthday time.Time
}

// AllClient select 
func AllClient() ([]Client, error) {
	rows, err := DB.Query("SELECT id, first_name, last_name, ci, birthday FROM clients LIMIT 10")
	if err != nil {
		return nil, err
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