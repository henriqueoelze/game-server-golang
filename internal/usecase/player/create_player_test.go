package player_test

import (
	"errors"
	"game-server-golang/internal/domain"
	"game-server-golang/internal/test/mocks"
	"game-server-golang/internal/usecase/player"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreatePlayer(t *testing.T) {
	t.Parallel()

	t.Run("should create player successfully", func(t *testing.T) {
		t.Parallel()

		// Arrange
		mockRepo := mocks.NewMockPlayerRepository(t)
		playerName := "New Player"
		playerLevel := 1
		mockRepo.EXPECT().
			CreatePlayer(mock.Anything).
			Run(func(player domain.Player) {
				assert.Equal(t, playerName, player.Name)
				assert.Equal(t, playerLevel, player.Level)
				assert.NotEmpty(t, player.PublicID)
			}).
			Return(nil)
		usecase := player.NewPlayerUsecase(mockRepo)

		// Act
		player, err := usecase.CreatePlayer()

		// Assert
		require.NoError(t, err)
		assert.Equal(t, playerName, player.Name)
		assert.Equal(t, playerLevel, player.Level)
		assert.NotEmpty(t, player.PublicID)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		t.Parallel()

		// Arrange
		mockRepo := mocks.NewMockPlayerRepository(t)
		expectedError := errors.New("database error")
		mockRepo.EXPECT().
			CreatePlayer(mock.Anything).
			Return(expectedError)
		usecase := player.NewPlayerUsecase(mockRepo)

		// Act
		player, err := usecase.CreatePlayer()

		// Assert
		require.Error(t, err)
		assert.Equal(t, domain.Player{}, player)
		assert.ErrorContains(t, err, "error creating player")
	})
}
