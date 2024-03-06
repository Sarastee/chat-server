package chat

import (
	"context"
	"fmt"
)

// Delete ...
func (s *Service) Delete(ctx context.Context, chatID int64) error {
	if err := s.chatRepo.Delete(ctx, chatID); err != nil {
		return fmt.Errorf("failure while deleteing chat: %w", err)
	}

	return nil
}
