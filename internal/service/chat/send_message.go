package chat

import (
	"context"
	"fmt"

	"github.com/sarastee/chat-server/internal/model"
	"github.com/sarastee/chat-server/internal/service"
)

// SendMessage ...
func (s *Service) SendMessage(ctx context.Context, message model.Message) error {
	userInChat, err := s.chatRepo.IsUserInChat(ctx, message.ToChatID, message.FromUserID)
	if err != nil {
		return fmt.Errorf("failure while sending message: %w", err)
	}

	if !userInChat {
		return service.ErrMsgUserNotInTheChat
	}

	if err := s.messageRepo.Create(ctx, message); err != nil {
		return fmt.Errorf("failure while sending message: %w", err)
	}

	return nil
}
