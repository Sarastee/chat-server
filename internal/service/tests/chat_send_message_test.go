package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/sarastee/chat-server/internal/model"
	repoMocks "github.com/sarastee/chat-server/internal/repository/mocks"
	chatService "github.com/sarastee/chat-server/internal/service"
	"github.com/sarastee/chat-server/internal/service/chat"
	"github.com/sarastee/platform_common/pkg/db/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreate_SuccessSendMessage(t *testing.T) {
	ctx := context.Background()

	message := model.Message{
		FromUserID: gofakeit.Int64(),
		Text:       gofakeit.BeerName(),
		ToChatID:   1,
		SendTime:   time.Now(),
	}

	txManagerMock := mocks.NewTxManager(t)

	chatRepoMock := repoMocks.NewChatRepository(t)
	chatRepoMock.On("IsUserInChat", ctx, message.ToChatID, message.FromUserID).Return(true, nil).Once()

	userRepoMock := repoMocks.NewUserRepository(t)

	messageRepoMock := repoMocks.NewMessageRepository(t)
	messageRepoMock.On("Create", ctx, message).Return(nil).Once()

	service := chat.NewService(txManagerMock, chatRepoMock, userRepoMock, messageRepoMock)
	err := service.SendMessage(ctx, message)

	require.NoError(t, err)
}

func TestCreate_FailSendMessageUserNotInTheChat(t *testing.T) {
	ctx := context.Background()

	message := model.Message{
		FromUserID: gofakeit.Int64(),
		Text:       gofakeit.BeerName(),
		ToChatID:   1,
		SendTime:   time.Now(),
	}

	txManagerMock := mocks.NewTxManager(t)

	chatRepoMock := repoMocks.NewChatRepository(t)
	chatRepoMock.On("IsUserInChat", ctx, message.ToChatID, message.FromUserID).Return(false, nil).Once()

	userRepoMock := repoMocks.NewUserRepository(t)
	messageRepoMock := repoMocks.NewMessageRepository(t)

	service := chat.NewService(txManagerMock, chatRepoMock, userRepoMock, messageRepoMock)
	err := service.SendMessage(ctx, message)

	require.Error(t, err)
	require.ErrorIs(t, chatService.ErrMsgUserNotInTheChat, err)
}

func TestCreate_FailSendMessageErrorInTimeCheckUserInTheChat(t *testing.T) {
	ctx := context.Background()
	checkErr := errors.New("some error")

	message := model.Message{
		FromUserID: gofakeit.Int64(),
		Text:       gofakeit.BeerName(),
		ToChatID:   1,
		SendTime:   time.Now(),
	}

	txManagerMock := mocks.NewTxManager(t)

	chatRepoMock := repoMocks.NewChatRepository(t)
	chatRepoMock.On("IsUserInChat", ctx, message.ToChatID, message.FromUserID).Return(false, checkErr).Once()

	userRepoMock := repoMocks.NewUserRepository(t)
	messageRepoMock := repoMocks.NewMessageRepository(t)

	service := chat.NewService(txManagerMock, chatRepoMock, userRepoMock, messageRepoMock)
	err := service.SendMessage(ctx, message)

	require.Error(t, err)
}

func TestCreate_FailSendMessageErrorInTimeSaveMessage(t *testing.T) {
	ctx := context.Background()
	saveErr := errors.New("some error")

	message := model.Message{
		FromUserID: gofakeit.Int64(),
		Text:       gofakeit.BeerName(),
		ToChatID:   1,
		SendTime:   time.Now(),
	}

	txManagerMock := mocks.NewTxManager(t)

	chatRepoMock := repoMocks.NewChatRepository(t)
	chatRepoMock.On("IsUserInChat", ctx, message.ToChatID, message.FromUserID).Return(true, nil).Once()

	userRepoMock := repoMocks.NewUserRepository(t)

	messageRepoMock := repoMocks.NewMessageRepository(t)
	messageRepoMock.On("Create", ctx, message).Return(saveErr).Once()

	service := chat.NewService(txManagerMock, chatRepoMock, userRepoMock, messageRepoMock)
	err := service.SendMessage(ctx, message)

	require.Error(t, err)
}
