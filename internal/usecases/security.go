package usecases

type SecurityUsecase interface {
	Encrypt(message string) (encryptedString string, err error)
	Decrypt(message string) (decryptedString string, err error)
}
