package migration

import (
	"github.com/golangmalaga/golangmalaga/configuration"
	"github.com/golangmalaga/golangmalaga/models"
)

//Migrate permite crear las tablas en la base de datos
func Migrate()  {
	database := configuration.GetConnection()
	defer database.Close()
	database.CreateTable(&models.User{})
}
