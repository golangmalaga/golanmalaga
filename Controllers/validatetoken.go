package Controllers

import (
	"net/http"
	"github.com/golangmalaga/golangmalaga/Models"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/golangmalaga/golangmalaga/Commons"
	"context"
)

//ValidateToken validar el token del cliente
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)  {
	var m Models.Message
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&Models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return Commons.PublicKey, nil
		},
	)
	if err != nil {
		m.Code = http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				m.Message = "Su token han expirado"
				Commons.DisplayMessage(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Message = "Las firma del foken no coincide"
				Commons.DisplayMessage(w, m)
				return
			default:
				m.Message = "Su token no es válido"
				Commons.DisplayMessage(w, m)
				return

			}

		}
	}
	if token.Valid {
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*Models.Claim).User)
		next(w, r.WithContext(ctx))
	}else {
		m.Code = http.StatusUnauthorized
		m.Message = "Su token no es válido"
		Commons.DisplayMessage(w, m)
	}
}
