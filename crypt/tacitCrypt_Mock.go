package crypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type TacitCryptMock struct {
	HasGeneratePasswordError bool
}


func (crypt *TacitCryptMock) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if crypt.HasGeneratePasswordError {
		return hash, fmt.Errorf("___GENERIC_CRYPT_ERROR___")
	} else {
		return hash, err
	}
}

func (crypt *TacitCryptMock) CompareHashAndPassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err
}
