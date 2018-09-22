package main

import (
	"math/big"
	"encoding/json"
	// "fmt"
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"crypto/rsa"
	"encoding/base64"

	"tacit-api/crypt"
	"tacit-api/db"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
)

type webUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type env struct {
	ourDB    db.TacitDB
	logger   logrus.FieldLogger
	ourCrypt crypt.TacitCrypt
}

func (e *env) doCreateUser(c *gin.Context) {
	ctx := &realHttpContext{ginCtx: c}
	createUser(ctx, e.ourDB, e.ourCrypt, e.logger)
}

func (e *env) doLogin(c *gin.Context) {
	ctx := &realHttpContext{ginCtx: c}
	login(ctx, e.ourDB, e.ourCrypt, e.logger)
}

func (e *env) doCreatePost(c *gin.Context) {
	ctx := &realHttpContext{ginCtx: c}
	createPost(ctx, e.ourDB, e.logger)
}

func (e *env) doListPosts(c *gin.Context) {
	ctx := &realHttpContext{ginCtx: c}
	listPosts(ctx, e.ourDB, e.logger)
}

func main() {
	aLogger := logrus.New()
	aLogger.Info("Tacit-api has started")
	// defaultHost := "localhost"
	// defaultPort := "5432"
	dbUser := os.Getenv("DB_USER")
	if len(dbUser) == 0 {
		dbUser = "gorm"
	}
	defaultDb := "tacit_db"
	dbPassword := os.Getenv("DB_PASSWORD")
	if len(dbPassword) == 0 {
		dbPassword = "@"
	}
	dbPort := os.Getenv("DB_PORT")
	if len(dbPort) == 0 {
		dbPort = "3306"
	}

	connectionString := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:" + dbPort + ")/" + defaultDb + "?charset=utf8&parseTime=True&loc=Local"
	// connectionString := "host="+defaultHost+" port="+defaultPort+" user="+defaultUser+" dbname="+defaultDb+" sslmode=disable"
	dbHandle, err := gorm.Open("mysql", connectionString) // TODO:: enable ssl
	defer dbHandle.Close()

	if err != nil {
		aLogger.Errorln("There was an error opeing the db: ", err)
		// TODO :: should exit right away
	}

	aRealTacitDB := &db.RealTacitDB{
		GormDB: dbHandle,
	}

	// db.RunMigration(aRealTacitDB)

	r := gin.Default()

	r.Use(FUCKITIWRITEMYOWN())

	r.GET("/ping", func(c *gin.Context) {
		response := c.MustGet("thing").(string)
		c.JSON(200, gin.H{"message": "pong: " + response})
	})
	anEnv := &env{ourDB: aRealTacitDB, logger: aLogger, ourCrypt: &crypt.RealTacitCrypt{}}

	r.POST("/user", anEnv.doCreateUser)

	r.POST("/login", anEnv.doLogin)

	r.POST("/note", anEnv.doCreatePost)

	r.GET("/note", anEnv.doListPosts)

	r.Run()
}

func FUCKITIWRITEMYOWN() gin.HandlerFunc {
	fmt.Println("installing a middleware")
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			fmt.Println("what is the kid of the token: ", token.Header["kid"])
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
		}
		// Don't forget to validate the alg is what you expect:
		// token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// 	// since we only use the one private key to sign the tokens,
		// 	// we also only use its public counter part to verify
		// 	return verifyKey, nil
		// })
		c.Set("thing", tokenString)
		fmt.Printf("IN A MIDDLEWARE BITCH with: '%v'\n", token)
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

type Auth0Key struct  {
	Alg string
	Kty string
	Use string
	X5c []string
	N string
	E string
	Kid string
	X5t string
}
