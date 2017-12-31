package Commons

import (
	"net/http"
	"github.com/golangmalaga/golangmalaga/Models"
	"encoding/json"
	"github.com/labstack/gommon/log"
)

//DisplayMessage devuelve un mensaje al cliente
func DisplayMessage(w http.ResponseWriter, m Models.Message)  {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
