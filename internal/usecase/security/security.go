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

func NewSecurityUsecase() usecase.SecurityUsecase {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return &SecurityUsecaseImpl{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}
}
