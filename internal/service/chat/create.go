package chat

import (
	"context"
	"fmt"
)

// Create ...
func (s *Service) Create(ctx context.Context, userIDs []int64) (int64, error) {
	var chatID int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		chatID, txErr = s.chatRepo.Create(ctx)
		if txErr != nil {
			return fmt.Errorf("failure while creating chat: %w", txErr)
		}

		txErr = s.userRepo.CreateMass(ctx, userIDs)
		if txErr != nil {
			return fmt.Errorf("failure while creating users: %w", txErr)
		}

		txErr = s.chatRepo.LinkChatAndUsers(ctx, chatID, userIDs)
		if txErr != nil {
			return fmt.Errorf("failure while connecting users with chat %w", txErr)
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
