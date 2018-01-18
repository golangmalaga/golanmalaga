package configuration

import (
	"os"
	"github.com/labstack/gommon/log"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"fmt"
)

type configuration struct {
	Server string
	User string
	Password string
	Database string
}

func getConfiguration() configuration {
	var c configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// GetConnection obtiene una conexion a la bd postgreSQL

func GetConnection() *gorm.DB {
	c := getConfiguration()
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", c.Server, c.User, c.Database, c.Password)
	database, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return database
}