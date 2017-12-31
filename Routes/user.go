package Routes

import (
	"github.com/gorilla/mux"
	"github.com/golangmalaga/golangmalaga/Controllers"
	"github.com/urfave/negroni"
)



// SetUserRouter rutas para el registro de usuario
func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", Controllers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			// negroni.HandlerFunc(controllers.ValidateToken)
			negroni.Wrap(subRouter),
		),
	)
}