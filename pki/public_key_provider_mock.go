package pki

import (
	"crypto/rsa"
	"fmt"
)

type PublicKeyProviderMock struct {
	Keys map[string]*rsa.PublicKey
}

func (this *PublicKeyProviderMock) GetPublicKey(kid string) (*rsa.PublicKey, error) {

	if key := this.Keys[kid]; key != nil {
		return key, nil
	} else {
		return nil, fmt.Errorf("Error occured fetching kid in test code")
	}

}
