package middleware

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtValidation() gin.HandlerFunc {
	fmt.Println("installing a middleware")
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// fmt.Println("what is the kid of the token: ", token.Header["kid"])
			resp, err := http.Get("https://tacit.auth0.com/.well-known/jwks.json")
			if err != nil {
				fmt.Printf("Our public key is unavailable becase: %v\n", err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error readying auth0keys response: ", err)
			}

			var keys Auth0PublicKeys
			err = json.Unmarshal(body, &keys)
			if err != nil {
				fmt.Println("Error unmarsaling auth0keys", err)
			}
			if token.Header["kid"] == keys.Keys[0].Kid {
				fmt.Println("whats the big deal?!")
			} else {
				fmt.Println("well obviously, this fails!")
			}
			// fmt.Printf("Here is my key: '%v'\n", keys.Keys[0].N)

			data := readKeyBytes(keys.Keys[0].N)
			// fmt.Printf("Lets see some data %q\n", data)

			bigN := new(big.Int)
			bigN.SetBytes(data)
			// fmt.Printf("Here is a big.Int: %v\n", bigN)

			// bigN, worked := bigN.SetString(keys.Keys[0].N, 0)
			// if !worked {
			// 	fmt.Println("Cannot convert N to bin.Int")
			// }
			data = readKeyBytes(keys.Keys[0].E)
			bigE := new(big.Int)
			bigE.SetBytes(data)
			// fmt.Printf("Here is a big.Int: %v\n", bigE)
			intE := int(bigE.Int64())
			if err != nil {
				fmt.Println("error parsing e: ", err)
			}
			return &rsa.PublicKey{N: bigN, E: intE}, nil
		})
		if err != nil {
			fmt.Println("There was an issue parsing the JWT: ", err)
		} else {
			fmt.Println("Here is an explicit success!!!!")
			c.Set("authed", true)
		}
		// Don't forget to validate the alg is what you expect:
		// token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// 	// since we only use the one private key to sign the tokens,
		// 	// we also only use its public counter part to verify
		// 	return verifyKey, nil
		// })
	}
}

func readKeyBytes(keyPart string) []byte {

	data, err := base64.RawURLEncoding.DecodeString(keyPart)
	if err != nil {
		fmt.Println("error:", err)
	}

	if len(data)%2 != 0 {
		data = append([]byte{'\x00'}, data...)
	}
	return data
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
