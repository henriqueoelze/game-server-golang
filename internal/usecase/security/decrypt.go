package security

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
)

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
