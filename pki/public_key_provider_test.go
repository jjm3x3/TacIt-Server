package pki

import (
	"encoding/json"
	"fmt"
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
