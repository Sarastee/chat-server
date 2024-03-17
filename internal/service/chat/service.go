package chat

import (
	"github.com/sarastee/chat-server/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// Service ...
type Service struct {
	txManager   db.TxManager
	chatRepo    repository.ChatRepository
	userRepo    repository.UserRepository
	messageRepo repository.MessageRepository
}

// NewService ...
func NewService(
	txManager db.TxManager,
	chatRepo repository.ChatRepository,
	userRepo repository.UserRepository,
	messageRepo repository.MessageRepository,
) *Service {
	return &Service{
		txManager:   txManager,
		chatRepo:    chatRepo,
		userRepo:    userRepo,
		messageRepo: messageRepo,
	}
}
