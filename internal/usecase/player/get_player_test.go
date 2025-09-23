package player_test

import (
	"errors"
	"game-server-golang/internal/domain"
	mocks "game-server-golang/internal/test/mock"
	"game-server-golang/internal/usecase/player"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPlayer(t *testing.T) {
	t.Parallel()

	t.Run("should get player successfully", func(t *testing.T) {
		t.Parallel()

		// Arrange
		mockRepo := mocks.NewMockPlayerRepository(t)
		playerID := uuid.New()
		expextedPlayer := &domain.Player{
			PublicID: playerID,
			Name:     "Existing Player",
			Level:    5,
		}
		mockRepo.EXPECT().
			GetPlayer(playerID).
			Return(expextedPlayer, nil)
		usecase := player.NewPlayerUsecase(mockRepo)

		// Act
		player, err := usecase.GetPlayer(playerID)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expextedPlayer, player)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		t.Parallel()

		// Arrange
		mockRepo := mocks.NewMockPlayerRepository(t)
		playerID := uuid.New()
		expectedError := errors.New("database error")
		mockRepo.EXPECT().
			GetPlayer(playerID).
			Return(nil, expectedError)
		usecase := player.NewPlayerUsecase(mockRepo)

		// Act
		player, err := usecase.GetPlayer(playerID)

		// Assert
		require.Error(t, err)
		assert.Nil(t, player)
	})
}
