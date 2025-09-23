package security_test

import (
	"game-server-golang/internal/usecase/security"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	t.Parallel()

	t.Run("Success_EncryptsMessage", func(t *testing.T) {
		t.Parallel()

		// Arrange
		usecase := security.NewSecurityUsecase()
		playerID := uuid.New().String()

		// Act
		encrypted, err := usecase.Encrypt(playerID)

		// Assert
		require.NoError(t, err)
		assert.NotEmpty(t, encrypted)
		assert.NotEqual(t, playerID, encrypted)
	})
}
