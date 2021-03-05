package lib

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config connection
type Config struct {
	Host string
	Name string
	User string
	Password string
	Port string
}

// NewConfig connection
func NewConfig() *Config {
	load := godotenv.Load("lib/.env")

	if load != nil {
		panic(load)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	return &Config {
		Host: dbHost,
		Name: dbName,
		User: dbUser,
		Password: dbPwd,
		Port: dbPort,
	}
}

// ConnectionString postgresql driver
func (c *Config) ConnectionString() (conn string) {
	conn = fmt.Sprintf("host=%s dbname=%s sslmode=disable user=%s password=%s", c.Host, c.Name, c.User, c.Password)
	return 
}