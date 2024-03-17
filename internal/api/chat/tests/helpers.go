package tests

import (
	"time"

	"github.com/sarastee/chat-server/internal/model"
	"github.com/sarastee/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func sendMessageRequestWithSetup(fromUserID int64, toChatID int64, onTime time.Time) *chat_v1.SendMessageRequest {
	return &chat_v1.SendMessageRequest{
		From:      fromUserID,
		Text:      "msg",
		ToChatId:  toChatID,
		Timestamp: timestamppb.New(onTime),
	}
}

func messageWithSetup(fromUserID int64, toChatID int64, onTime time.Time) *model.Message {
	return &model.Message{
		FromUserID: fromUserID,
		Text:       "msg",
		ToChatID:   toChatID,
		SendTime:   onTime,
	}
}
