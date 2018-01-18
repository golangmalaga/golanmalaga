package routes

import (
	"github.com/gorilla/mux"
	"github.com/golangmalaga/golangmalaga/controllers"
)

//SetLoginRouter ruta para el login
func SetLoginRouter(router *mux.Router)  {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
