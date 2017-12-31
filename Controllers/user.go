package Controllers

import (
	"github.com/golangmalaga/golangmalaga/Models"
	"crypto/sha256"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/golangmalaga/golangmalaga/Configuration"
	"log"
	"github.com/golangmalaga/golangmalaga/Commons"
	"crypto/md5"
)

var (
	user = Models.User{}
	m = Models.Message{}
	c = sha256.Sum256([]byte(user.Password))
	password = fmt.Sprintf("%x", c)
)

// Login es el controlador de login
func Login(w http.ResponseWriter, r *http.Request)  {
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	database := Configuration.GetConnection()
	defer database.Close()

	database.Where("email = $1 and password = $2", user.Email, password).First(&user)
	log.Println(user.ID, password)
	if user.ID > 0 {
		user.Password = ""
		token :=  Commons.GenerateJWT(user)
		j, err := json.Marshal(Models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error al convertir el token a json: %s", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}else {
		m = Models.Message{
			Message: "Usuario o clave no válido",
			Code: http.StatusUnauthorized,
		}
		Commons.DisplayMessage(w, m)
	}

}

// UserCreate permite registrar un Usuario
func  UserCreate(w http.ResponseWriter, r *http.Request)  {
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a registrarse: %s", err)
		m.Code = http.StatusBadRequest
		Commons.DisplayMessage(w, m)
		return
	}
	if user.Password != user.ConfirmPassword {
		m.Message = "Las contraseña no coinciden"
		m.Code = http.StatusBadRequest
		Commons.DisplayMessage(w, m)
		return
	}
	user.Password = password
	avatarmd5 :=  md5.Sum([]byte(user.Password))
	avatarstr := fmt.Sprintf("%x", avatarmd5)
	user.Avatar = "https://gravatar.com/avatar/" + avatarstr + "?s=100"
	database := Configuration.GetConnection()
	defer  database.Close()
	err = database.Create(&user).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		m.Code = http.StatusBadRequest
		Commons.DisplayMessage(w, m)
		return
	}
	m.Message = "Usuario creado con éxito"
	m.Code = http.StatusCreated
	Commons.DisplayMessage(w, m)
}
