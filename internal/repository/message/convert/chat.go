package convert

import (
	serviceModel "github.com/sarastee/chat-server/internal/model"
	"github.com/sarastee/chat-server/internal/repository/message/model"
)

// ToMessageFromServiceMessage ...
func ToMessageFromServiceMessage(message *serviceModel.Message) *model.Message {
	return &model.Message{
		FromUserID: message.FromUserID,
		Text:       message.Text,
		ToChatID:   message.ToChatID,
		SendTime:   message.SendTime,
	}
}
