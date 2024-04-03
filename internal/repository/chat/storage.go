package chat

import (
	"github.com/sarastee/chat-server/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	chatTable     = "\"chat\""
	chatUserTable = "chat_user"

	userIDColumn = "user_id"
	chatIDColumn = "chat_id"
	idColumn     = "id"
)

var _ repository.ChatRepository = (*Repo)(nil)

// Repo ...
type Repo struct {
	db db.Client
}

// NewRepo ...
func NewRepo(dbClient db.Client) *Repo {
	return &Repo{
		db: dbClient,
	}
}
