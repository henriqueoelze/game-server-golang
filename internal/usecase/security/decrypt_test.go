package security_test

import (
	"game-server-golang/internal/usecase/security"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecrypt(t *testing.T) {
	t.Parallel()

	t.Run("Success_DecryptsMessage", func(t *testing.T) {
		t.Parallel()

		// Arrange
		usecase := security.NewSecurityUsecase()
		playerID := uuid.New().String()
		encrypted, err := usecase.Encrypt(playerID)
		require.NoError(t, err)

		// Act
		decrypted, err := usecase.Decrypt(encrypted)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, playerID, decrypted)
	})
}
