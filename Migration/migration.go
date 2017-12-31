package Migration

import (
	"github.com/golangmalaga/golangmalaga/Configuration"
	"github.com/golangmalaga/golangmalaga/Models"
)

//Migrate permite crear las tablas en la base de datos
func Migrate()  {
	database := Configuration.GetConnection()
	defer database.Close()
	database.CreateTable(&Models.User{})
}
