package pki

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublicKeyProviderHappyPath(t *testing.T) {
	// setup
	expectedId := "someId"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		keys := Auth0PublicKeys{}
		keys.Keys = append(keys.Keys, Auth0Key{Kid: expectedId})

		response, err := json.Marshal(keys)
		if err != nil {
			fmt.Println("There was an issue marshaling JSON: ", err)
		}
		rw.Write(response)
	}))

	defer server.Close()

	pkp := NewPublicKeyProvider(server.Client(), server.URL)

	// execution
	key, err := pkp.GetPublicKey(expectedId)

	// assertions
	if err != nil {
		t.Error("An error occured while trying to get the public key: ", err)
	}

	if key == nil {
		t.Error("We should get a non nil key in this test")
	}
}

func TestPublicKeyProviderFailsWhenTheServerStoringPublicKeysFails(t *testing.T) {
	// setup
	expectedId := "someId"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		panic("Server had a terrible failure")
	}))

	defer server.Close()

	pkp := NewPublicKeyProvider(server.Client(), server.URL)

	// execution
	key, err := pkp.GetPublicKey(expectedId)

	// assertions
	if err == nil {
		t.Error("An error is expected when the server storing public keys fails: ", err)
	}

	if key != nil {
		t.Error("We should get a nil key in this test")
	}
}
func TestPublicKeyProviderShouldReturnAnErrorIfItDoesntGetAnExpectedAuth0PublicKeyType(t *testing.T) {
	// setup
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("ok"))
	}))

	defer server.Close()

	pkp := NewPublicKeyProvider(server.Client(), server.URL)

	// execution
	key, err := pkp.GetPublicKey("someID")

	// assertions
	if err == nil {
		t.Error("An error is expected when an unknown id is passed into the get key method: ", err)
	}

	if key != nil {
		t.Error("We should get a nil key in this test")
	}
}

func TestPublicKeyProviderShouldReturnAnErrorIfItDoesntFindAKeyWithTheProvidedId(t *testing.T) {
	// setup
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		keys := Auth0PublicKeys{}
		keys.Keys = append(keys.Keys, Auth0Key{})

		response, err := json.Marshal(keys)
		if err != nil {
			fmt.Println("There was an issue marshaling JSON: ", err)
		}
		rw.Write(response)
	}))

	defer server.Close()

	pkp := NewPublicKeyProvider(server.Client(), server.URL)

	// execution
	key, err := pkp.GetPublicKey("someID")

	// assertions
	if err == nil {
		t.Error("An error is expected when an unknown id is passed into the get key method: ", err)
	}

	if key != nil {
		t.Error("We should get a nil key in this test")
	}
}

func TestPublicKeyProviderReturnsTheRightKey(t *testing.T) {
	// setup
	expectedId := "someId"
	expectedKey := GenRsaPair()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		keys := Auth0PublicKeys{}

		nString := base64.RawURLEncoding.EncodeToString(expectedKey.N.Bytes())
		// using the big int strategty to encode E 
		bigE := new(big.Int)
		bigE.SetInt64(int64(expectedKey.E))
		eString := base64.RawURLEncoding.EncodeToString(bigE.Bytes())
		keys.Keys = append(keys.Keys, Auth0Key{Kid: "unexpectedId"})
		keys.Keys = append(keys.Keys, Auth0Key{Kid: expectedId, N: nString, E: eString})

		response, err := json.Marshal(keys)
		if err != nil {
			fmt.Println("There was an issue marshaling JSON: ", err)
		}
		rw.Write(response)
	}))

	defer server.Close()

	pkp := NewPublicKeyProvider(server.Client(), server.URL)

	// execution
	key, err := pkp.GetPublicKey(expectedId)

	// assertions
	if err != nil {
		t.Error("There should be no error when trying to get the correct key: ", err)
	}

	if key.E != expectedKey.PublicKey.E {
		t.Errorf("Expected key.E: '%v', got key.E '%v'", expectedKey.PublicKey.E, key.E)
	}

	if key.N.String() != expectedKey.PublicKey.N.String() {
		t.Errorf("Expected key.N: '%v', got key.N '%v'", expectedKey.PublicKey.N, key.N)
	}
}
