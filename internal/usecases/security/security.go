package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"game-server-golang/internal/base"
	"game-server-golang/internal/usecases"
)

var _ = usecases.SecurityUsecase(&SecurityUsecaseImpl{})

type SecurityUsecaseImpl struct {
	base.BaseLogger
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewSecurityUsecase() usecases.SecurityUsecase {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return &SecurityUsecaseImpl{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}
}

func (s *SecurityUsecaseImpl) Decrypt(encryptedMessage string) (decryptedString string, err error) {
	data, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", err
	}
	decryptedBytes, err := s.privateKey.Decrypt(nil, []byte(data), &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return "", err
	}

	return string(decryptedBytes), nil
}

func (s *SecurityUsecaseImpl) Encrypt(message string) (encryptedString string, err error) {
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		s.publicKey,
		[]byte(message),
		nil,
	)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}
