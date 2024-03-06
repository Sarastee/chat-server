package chat

import (
	"github.com/sarastee/chat-server/internal/client/db"
	"github.com/sarastee/chat-server/internal/repository"
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
