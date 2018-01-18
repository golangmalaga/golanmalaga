package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"github.com/labstack/gommon/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/golangmalaga/golangmalaga/models"
)

var (
	privateKey *rsa.PrivateKey
	// PublicKey se usara para validar el token
	PublicKey *rsa.PublicKey
)

func init()  {
	privateBytes, err := ioutil.ReadFile("./Keys/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}
	publicBytes, err := ioutil.ReadFile("./Keys/public.rsa")
	if err != nil {
		log.Fatal("NO se pudo leer el archivo publico")
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo parsear la llave privada", err)
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo parsear la llave publica")
	}
}

// GenerateJWT genera el token para el cliente
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer: "Golangmalaga",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err :=  token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}
	return  result
}
