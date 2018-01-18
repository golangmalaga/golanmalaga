package controllers

import (
	"github.com/golangmalaga/golangmalaga/models"
	"crypto/sha256"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/golangmalaga/golangmalaga/configuration"
	"log"
	"github.com/golangmalaga/golangmalaga/commons"
	"crypto/md5"
)

var (
	user = models.User{}
	m = models.Message{}
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
	database := configuration.GetConnection()
	defer database.Close()

	database.Where("email = $1 and password = $2", user.Email, password).First(&user)
	log.Println(user.ID, password)
	if user.ID > 0 {
		user.Password = ""
		token :=  commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error al convertir el token a json: %s", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}else {
		m = models.Message{
			Message: "Usuario o clave no válido",
			Code: http.StatusUnauthorized,
		}
		commons.DisplayMessage(w, m)
	}

}

// UserCreate permite registrar un Usuario
func  UserCreate(w http.ResponseWriter, r *http.Request)  {
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a registrarse: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	if user.Password != user.ConfirmPassword {
		m.Message = "Las contraseña no coinciden"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	user.Password = password
	avatarmd5 :=  md5.Sum([]byte(user.Password))
	avatarstr := fmt.Sprintf("%x", avatarmd5)
	user.Avatar = "https://gravatar.com/avatar/" + avatarstr + "?s=100"
	database := configuration.GetConnection()
	defer  database.Close()
	err = database.Create(&user).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al crear el registro: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	m.Message = "Usuario creado con éxito"
	m.Code = http.StatusCreated
	commons.DisplayMessage(w, m)
}
