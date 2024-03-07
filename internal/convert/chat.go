package convert

import (
	"github.com/sarastee/chat-server/internal/model"
	desc "github.com/sarastee/chat-server/pkg/chat_v1"
)

// ToMessageFromDesc ..
func ToMessageFromDesc(request *desc.SendMessageRequest) *model.Message {
	return &model.Message{
		FromUserID: request.From,
		Text:       request.Text,
		ToChatID:   request.ToChatId,
		SendTime:   request.Timestamp.AsTime(),
	}
}
