package main

import (
	"flag"
	"log"
	"github.com/golangmalaga/golangmalaga/Migration"
)

func main()  {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la Base de datos")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración a la base de datos...")
		Migration.Migrate()
		log.Println("Se finalizó la migración.")
	}
}