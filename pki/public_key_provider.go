package pki

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

type PublicKeyProvider interface {
	GetPublicKey(string) (*rsa.PublicKey, error)
}

type RealPublicKeyProvider struct {
	url    string
	client *http.Client
}

func NewPublicKeyProvider(client *http.Client, url string) PublicKeyProvider {
	return &RealPublicKeyProvider{client: client, url: url}
}

func (this *RealPublicKeyProvider) GetPublicKey(keyId string) (*rsa.PublicKey, error) {
	resp, err := this.client.Get(this.url)
	if err != nil {
		return nil, fmt.Errorf("Our public key is unavailable becase: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO:: not sure how to test this or if I need to
		return nil, fmt.Errorf("Error reading auth0keys response: %v", err)
	}

	var keys Auth0PublicKeys
	err = json.Unmarshal(body, &keys)
	if err != nil {
		return nil, fmt.Errorf("Error unmarsaling auth0keys %v", err)
	}

	var validaingKey Auth0Key
	for _, aKey := range keys.Keys {
		if aKey.Kid == keyId {
			validaingKey = aKey
		}
	}

	if len(validaingKey.Kid) == 0 {
		return nil, fmt.Errorf("There was no matching key for kid: %v", keyId)
	}

	// TODO:: GET correct key
	data, err := readKeyBytes(keys.Keys[0].N)
	if err != nil {
		return nil, err
	}

	bigN := new(big.Int)
	bigN.SetBytes(data)

	// TODO:: GET correct key
	data, err = readKeyBytes(keys.Keys[0].E)
	if err != nil {
		return nil, err
	}

	bigE := new(big.Int)
	bigE.SetBytes(data)
	intE := int(bigE.Int64()) // safe since this should be a small number

	return &rsa.PublicKey{N: bigN, E: intE}, nil
}

func readKeyBytes(keyPart string) ([]byte, error) {

	data, err := base64.RawURLEncoding.DecodeString(keyPart)
	if err != nil {
		return nil, fmt.Errorf("Decoding Error: %v", err)
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
