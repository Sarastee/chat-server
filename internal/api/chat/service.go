package chat

import (
	"errors"

	"github.com/sarastee/chat-server/internal/service"
	"github.com/sarastee/chat-server/pkg/chat_v1"
)

const msgInternalError = "something went wrong, we are already working on it"

var errInternal = errors.New(msgInternalError)

// Implementation ..
type Implementation struct {
	chat_v1.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewImplementation ..
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
