package routes

import "github.com/gorilla/mux"

// InitRoutes inicias la rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	SetPublicRouter(router)

	return router
}
