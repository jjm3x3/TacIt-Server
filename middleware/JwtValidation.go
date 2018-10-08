package middleware

import (
	"fmt"

	tacitHttp "tacit-api/http"
	pki "tacit-api/pki"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

func JwtValidation(callContext tacitHttp.HttpContext, fieldLogger logrus.FieldLogger, publickcKeyProvider pki.PublicKeyProvider) {
	tokenString := callContext.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		kidString := token.Header["kid"]
		if kidString == nil {
			return nil, fmt.Errorf("Missing kid in jwt cannot validate")
		}
		tokenKid := kidString.(string)
		fieldLogger.Debugf("KID of token %v", tokenKid)

		return publickcKeyProvider.GetPublicKey(tokenKid)
	})
	if err != nil {
		fieldLogger.Info("There was an issue validating the JWT: ", err)
		return
	}

	fieldLogger.Info("Here is an explicit success!!!!")
	callContext.Set("authed", true)

	fieldLogger.Info("What are my claims?: ", token.Claims)

	// callContext.Next()
}
