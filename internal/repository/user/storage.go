package user

import (
	"github.com/sarastee/chat-server/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	userTable = "\"user\""

	idColumn = "id"
)

var _ repository.UserRepository = (*Repo)(nil)

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
