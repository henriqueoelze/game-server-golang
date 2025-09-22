package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func (s *SecurityUsecaseImpl) Encrypt(message string) (string, error) {
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		s.publicKey,
		[]byte(message),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("error encrypting message: %w", err)
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}
