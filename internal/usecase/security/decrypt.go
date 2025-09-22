package security

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

func (s *SecurityUsecaseImpl) Decrypt(encryptedMessage string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 message: %w", err)
	}

	decryptedBytes, err := s.privateKey.Decrypt(nil, data, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return "", fmt.Errorf("error decrypting message: %w", err)
	}

	return string(decryptedBytes), nil
}
