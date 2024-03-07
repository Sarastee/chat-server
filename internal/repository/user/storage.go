package user

import (
	"github.com/Masterminds/squirrel"
	"github.com/sarastee/chat-server/internal/client/db"
	"github.com/sarastee/chat-server/internal/repository"
)

const (
	userTable = "\"user\""

	idColumn = "id"
)

var _ repository.UserRepository = (*Repo)(nil)

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
