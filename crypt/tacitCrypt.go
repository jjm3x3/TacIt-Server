package crypt 

import (
	"golang.org/x/crypto/bcrypt"
)

type RealTacitCrypt struct {
}

type TacitCrypt interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

func (crypt *RealTacitCrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	return hash, err
}

func (crypt *RealTacitCrypt) CompareHashAndPassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err
}
