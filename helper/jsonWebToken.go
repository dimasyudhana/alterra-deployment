package helper

import (
	"log"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(phone string) string {
	var information = jwt.MapClaims{}
	information["phone"] = phone

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, information)

	resultToken, err := rawToken.SignedString([]byte("@secret99"))
	if err != nil {
		log.Println("error generating JWT", err.Error())
		return ""
	}

	return resultToken

}

func DecodeJWT(token *jwt.Token) string {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		phone := data["phone"].(string)

		return phone
	}
	return ""
}
