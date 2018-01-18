package Routes

import (
	"github.com/gorilla/mux"
	"github.com/golangmalaga/golangmalaga/Controllers"
)

//SetLoginRouter ruta para el login
func SetLoginRouter(router *mux.Router)  {
	router.HandleFunc("/api/login", Controllers.Login).Methods("POST")
}
