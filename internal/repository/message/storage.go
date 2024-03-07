package message

import (
	"github.com/Masterminds/squirrel"
	"github.com/sarastee/chat-server/internal/client/db"
	"github.com/sarastee/chat-server/internal/repository"
)

const (
	messageTable = "message"

	chatIDColumn     = "chat_id"
	fromUserIDColumn = "from_user_id"
	textColumn       = "text"
	sentAtColumn     = "sent_at"
	idColumn         = "id"

	messageChatIDFKConstraint     = "fk_chat_id"
	messageFromUserIDFKConstraint = "fk_from_user_id"
)

var _ repository.MessageRepository = (*Repo)(nil)

// Repo ...
type Repo struct {
	db db.Client
	sq squirrel.StatementBuilderType
}

// NewRepo ...
func NewRepo(dbClient db.Client) *Repo {
	return &Repo{
		db: dbClient,
	}
}
