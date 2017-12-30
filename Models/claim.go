package Models

import "github.com/dgrijalva/jwt-go"

//Claim token de los usuario
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
