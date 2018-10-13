package middleware

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	http "tacit-api/http"
	"tacit-api/mocks"
	pki "tacit-api/pki"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
)

func TestDoJwtValidationFailsWithBadJwt(t *testing.T) {

	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &http.HttpContextMock{
		GetHeaderResult: getHeaderAndBody(expectedAudience, expectedIssuer),
	}

	pkp := &pki.RealPublicKeyProvider{}

	mockLogger.EXPECT().Info("There was an issue validating the JWT: ", gomock.Any())

	// execution
	JwtValidation(c, mockLogger, pkp)

	//assertions
	if c.SetIsCalled {
		t.Errorf("Set should not be called since the JWT cannot be validated")
	}
}

func TestDoJwtValidationFailsWithoutAKID(t *testing.T) {

	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &http.HttpContextMock{
		GetHeaderResult: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.TCYt5XsITJX1CxPCT8yAV-TVkIEq_PbChOMqsLfRoPsnsgw5WEuts01mq-pQy7UJiN5mgRxD-WUcX16dUEMGlv50aqzpqh4Qktb3rk-BuQy72IFLOqV0G_zS245-kronKb78cPN25DGlcTwLtjPAYuNzVBAh4vGHSrQyHUdBBPM",
	}

	pkp := &pki.RealPublicKeyProvider{}

	mockLogger.EXPECT().Info("There was an issue validating the JWT: ", gomock.Any())

	// execution
	JwtValidation(c, mockLogger, pkp)

	//assertions
	if c.SetIsCalled {
		t.Errorf("Set should not be called since the JWT cannot be validated")
	}
}

func genRsaPair() *rsa.PrivateKey {

	reader := rand.Reader
	bitSize := 512

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		fmt.Println("Can't generate rsa key pair?")
	}
	return key
}

func getHeaderAndBody(aud, iss string) string {

	jwtHeaderJsonString := `{"alg":"RS256","typ":"JWT","kid":"M0E1MzQzMjM4RDEwNzI4RDE0NzE5QTE3RTlDNkU1NTc0QThGREE3MA"}`
	header := base64.RawURLEncoding.EncodeToString([]byte(jwtHeaderJsonString))
	jwtClaims := jwt.StandardClaims{
		Subject:  "1234567890",
		IssuedAt: 1516239022,
		Audience: aud,
		Issuer:   iss,
	}
	jwtBodyJsonString, err := json.Marshal(jwtClaims) //`{"name":"JohnDoe","admin":true}`
	if err != nil {
		fmt.Println("There was an issue marshalling the json: ", jwtClaims)
	}
	body := base64.RawURLEncoding.EncodeToString([]byte(jwtBodyJsonString))

	return header + "." + body
}

func genJwt(key *rsa.PrivateKey, aud, iss string) string {
	if key == nil {
		key = genRsaPair()
	}

	headerAndBody := getHeaderAndBody(aud, iss)
	hashedHeaderAndBody := sha256.Sum256([]byte(headerAndBody))
	sigBytes, err := key.Sign(rand.Reader, hashedHeaderAndBody[:], crypto.SHA256)
	if err != nil {
		fmt.Println("there was an error signing JWT: ", err)
	}
	sigBase64 := base64.RawURLEncoding.EncodeToString(sigBytes)

	return headerAndBody + "." + sigBase64
}

func TestDoJwtValidationHappyPath(t *testing.T) {

	//setup
	key := genRsaPair()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &http.HttpContextMock{
		GetHeaderResult: genJwt(key, expectedAudience, expectedIssuer),
	}

	pkp := &pki.PublicKeyProviderMock{Keys: make(map[string]*rsa.PublicKey)}
	pkp.Keys["M0E1MzQzMjM4RDEwNzI4RDE0NzE5QTE3RTlDNkU1NTc0QThGREE3MA"] = &key.PublicKey

	mockLogger.EXPECT().Debugf("KID of token %v", gomock.Any())
	mockLogger.EXPECT().Debug("JWT has been validated")

	// execution
	JwtValidation(c, mockLogger, pkp)

	if !c.SetIsCalled {
		t.Error("Set should have been called since we expect this JWT to be valid")
	}

	mockLogger.EXPECT().Debugf("KID of token %v", gomock.Any())

	// execution
	JwtValidation(c, mockLogger, pkp)

}
