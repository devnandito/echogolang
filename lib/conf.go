package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Config connection
type Config struct {
	Host string
	Name string
	User string
	Password string
	Port string
}

// GetEnv load env
func GetEnv() (host, name, user, pwd, port string) {
	err := godotenv.Load("lib/.env")
	if err != nil {
	 	log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbPort := os.Getenv("DB_PORT")
	return dbHost, dbName, dbUser, dbPwd, dbPort
 }

// NewConfig connection
func NewConfig() *Config {
	host, name, user, pwd, port  := GetEnv()
	return &Config {
		Host: host,
		Name: name,
		User: user,
		Password: pwd,
		Port: port,
	}
}

// DsnString postgresql driver
func (c *Config) DsnString() (conn *sql.DB) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable user=%s password=%s port=%s", c.Host, c.Name, c.User, c.Password, c.Port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db
}