package main

import (
	"flag"
	"log"
	"github.com/golangmalaga/golangmalaga/Migration"
	"github.com/golangmalaga/golangmalaga/Routes"
	"github.com/urfave/negroni"
	"net/http"
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
	// Inicia las rutas
	router := Routes.InitRoutes()

	//Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr: ":8080",
		Handler: n,
	}
	log.Println("Iniciando el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa")
}