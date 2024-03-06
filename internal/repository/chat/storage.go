package chat

import (
	"github.com/Masterminds/squirrel"
	"github.com/sarastee/chat-server/internal/client/db"
	"github.com/sarastee/chat-server/internal/repository"
)

const (
	chatTable     = "chat"
	chatUserTable = "chat_user"

	userIDColumn = "user_id"
	chatIDColumn = "chat_id"
	idColumn     = "id"
)

var _ repository.ChatRepository = (*Repo)(nil)

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
