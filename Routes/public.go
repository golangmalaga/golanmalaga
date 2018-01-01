package Routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

// SetPublicRouter expone los archivos estáticos
func SetPublicRouter(router *mux.Router)  {
	router.Handle("/", http.FileServer(http.Dir("./public")))
}
