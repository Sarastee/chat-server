package tests

import (
	"context"
	"errors"
	"testing"

	repoMocks "github.com/sarastee/chat-server/internal/repository/mocks"
	"github.com/sarastee/chat-server/internal/service/chat"
	"github.com/sarastee/platform_common/pkg/db/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreate_SuccessDeleteChat(t *testing.T) {
	ctx := context.Background()
	var chatID int64 = 1

	txManagerMock := mocks.NewTxManager(t)

	chatRepoMock := repoMocks.NewChatRepository(t)
	chatRepoMock.On("Delete", ctx, chatID).Return(nil).Once()

	userRepoMock := repoMocks.NewUserRepository(t)
	messageRepoMock := repoMocks.NewMessageRepository(t)

	service := chat.NewService(txManagerMock, chatRepoMock, userRepoMock, messageRepoMock)
	err := service.Delete(ctx, chatID)

	require.NoError(t, err)
}

func TestCreate_FailDeleteChat(t *testing.T) {
	ctx := context.Background()
	var chatID int64 = 1
	deleteError := errors.New("some error")

	txManagerMock := mocks.NewTxManager(t)

	chatRepoMock := repoMocks.NewChatRepository(t)
	chatRepoMock.On("Delete", ctx, chatID).Return(deleteError).Once()

	userRepoMock := repoMocks.NewUserRepository(t)
	messageRepoMock := repoMocks.NewMessageRepository(t)

	service := chat.NewService(txManagerMock, chatRepoMock, userRepoMock, messageRepoMock)
	err := service.Delete(ctx, chatID)

	require.Error(t, err)
}
