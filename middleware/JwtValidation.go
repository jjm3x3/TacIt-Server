package middleware

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	tacitHttp "tacit-api/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JwtValidation(fieldLogger logrus.FieldLogger) gin.HandlerFunc {
	fieldLogger.Info("Installing JWT Middleware")
	return func(c *gin.Context) {
		doJwtValidation(fieldLogger, tacitHttp.NewContext(c))
	}
}

func doJwtValidation(fieldLogger logrus.FieldLogger, callContext tacitHttp.HttpContext) {
	tokenString := callContext.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		tokenKid := token.Header["kid"]
		fieldLogger.Debugf("KID of token %v", tokenKid)

		resp, err := http.Get("https://tacit.auth0.com/.well-known/jwks.json")
		if err != nil {
			return nil, fmt.Errorf("Our public key is unavailable becase: %v\n", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Error reading auth0keys response: ", err)
		}

		var keys Auth0PublicKeys
		err = json.Unmarshal(body, &keys)
		if err != nil {
			return nil, fmt.Errorf("Error unmarsaling auth0keys", err)
		}

		var validaingKey Auth0Key
		for _, aKey := range keys.Keys {
			if aKey.Kid == tokenKid {
				validaingKey = aKey
			}
		}

		if len(validaingKey.Kid) == 0 {
			return nil, fmt.Errorf("There was no matching key for kid: ", tokenKid)
		}

		data, err := readKeyBytes(keys.Keys[0].N)
		if err != nil {
			return nil, err
		}

		bigN := new(big.Int)
		bigN.SetBytes(data)

		data, err = readKeyBytes(keys.Keys[0].E)
		if err != nil {
			return nil, err
		}

		bigE := new(big.Int)
		bigE.SetBytes(data)
		intE := int(bigE.Int64()) // safe since this should be a small number

		return &rsa.PublicKey{N: bigN, E: intE}, nil
	})
	if err != nil {
		fmt.Println("There was an issue parsing the JWT: ", err)
	} else {
		fmt.Println("Here is an explicit success!!!!")
		fmt.Println("What are my claims?: ", token.Claims)
		callContext.Set("authed", true)
	}
	// Don't forget to validate the alg is what you expect:
	// token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
	// 	// since we only use the one private key to sign the tokens,
	// 	// we also only use its public counter part to verify
	// 	return verifyKey, nil
	// })
}

func readKeyBytes(keyPart string) ([]byte, error) {

	data, err := base64.RawURLEncoding.DecodeString(keyPart)
	if err != nil {
		return nil, fmt.Errorf("Decoding Error:", err)
	}

	if len(data)%2 != 0 {
		data = append([]byte{'\x00'}, data...)
	}
	return data, nil
}

type Auth0PublicKeys struct {
	Keys []Auth0Key
}

type Auth0Key struct {
	Alg string
	Kty string
	Use string
	X5c []string
	N   string
	E   string
	Kid string
	X5t string
}
