package security

import (
	"crypto/rand"
	"crypto/rsa"

	"game-server-golang/internal/usecase"
)

var _ = usecase.SecurityUsecase(&SecurityUsecaseImpl{})

type SecurityUsecaseImpl struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewSecurityUsecase() *SecurityUsecaseImpl {
	numberOfBits := 2048
	privateKey, err := rsa.GenerateKey(rand.Reader, numberOfBits)

	if err != nil {
		panic(err)
	}

	return &SecurityUsecaseImpl{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}
}
