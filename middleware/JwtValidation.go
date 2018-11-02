package middleware

import (
	"fmt"

	tacitHttp "tacit-api/http"
	pki "tacit-api/pki"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var expectedAudience = "http://tacit-dev.tacitapp.io"
var expectedIssuer = "https://tacit.auth0.com/"

func JwtValidation(callContext tacitHttp.HttpContext, fieldLogger logrus.FieldLogger, publickcKeyProvider pki.PublicKeyProvider) {
	tokenString := callContext.GetHeader("Authorization")
	_, err := jwt.ParseWithClaims(tokenString, &Auth0Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		kidString := token.Header["kid"]
		if kidString == nil {
			return nil, fmt.Errorf("Missing kid in jwt cannot validate")
		}
		tokenKid := kidString.(string)
		fieldLogger.Debugf("KID of token %v", tokenKid)

		if claims, ok := token.Claims.(*Auth0Claims); ok {
			err := claims.StandardClaims.Valid()
			if err != nil {
				return nil, err
			}
			if ok := claims.VerifyAudience(expectedAudience); !ok {
				return nil, fmt.Errorf("Audience for token is unexpected")
			}
			if ok := claims.StandardClaims.VerifyIssuer(expectedIssuer, true); !ok {
				return nil, fmt.Errorf("Issuer for token is unexpected")
			}

			fmt.Println("Here are some claims: ", claims.Scope)
		}

		return publickcKeyProvider.GetPublicKey(tokenKid)
	})
	if err != nil {
		fieldLogger.Info("There was an issue validating the JWT: ", err)
		return
	}

	fieldLogger.Debug("JWT has been validated")
	callContext.Set("authed", true)

	// fieldLogger.Info("What are my claims?: ", token.Claims)

	// callContext.Next()
}

type Auth0Claims struct {
	Scope    string   `json:scope`
	Audience []string `json:"aud,omitempty"`
	jwt.StandardClaims
}

// need to recreate since we need "override" StandardClaims.Audience
func (c *Auth0Claims) VerifyAudience(cmp string) bool {
	for _, aud := range c.Audience {
		if verifyAud(aud, cmp, true) {
			return true
		}
	}
	return false
}
