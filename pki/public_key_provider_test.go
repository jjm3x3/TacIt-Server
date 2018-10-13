package pki

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublicKeyProviderHappyPath(t *testing.T) {
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

	key, err := pkp.GetPublicKey(expectedId)

	if err != nil {
		t.Error("An error occured while trying to get the public key: ", err)
	}

	if key == nil {
		t.Error("We should get a non nil key in this test")
	}

	fmt.Println("here is the final key result: ", key)
	t.Fail()
}

func TestPublicKeyProviderShouldReturnAnErrorIfItDoesntGetAnExpectedAuth0PublicKeyType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("ok"))
	}))

	defer server.Close()

	pkp := NewPublicKeyProvider(server.Client(), server.URL)

	key, err := pkp.GetPublicKey("someID")

	if err == nil {
		t.Error("An error is expected when an unknown id is passed into the get key method: ", err)
	}

	if key != nil {
		t.Error("We should get a nil key in this test")
	}
}

func TestPublicKeyProviderShouldReturnAnErrorIfItDoesntFindAKeyWithTheProvidedId(t *testing.T) {
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

	key, err := pkp.GetPublicKey("someID")

	if err == nil {
		t.Error("An error is expected when an unknown id is passed into the get key method: ", err)
	}

	if key != nil {
		t.Error("We should get a nil key in this test")
	}
}
