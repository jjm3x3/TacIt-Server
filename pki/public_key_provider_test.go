package pki

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHappyPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("ok"))
	}))

	defer server.Close()

	pkp := NewPublicKeyProvider(server.Client(), server.URL)

	key, err := pkp.GetPublicKey("someID")

	if err != nil {
		t.Error("An error occured while trying to get the public key: ", err)
	}

	if key == nil {
		t.Error("We should get a non nil key in this test")
	}
}
