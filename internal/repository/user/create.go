package user

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/sarastee/chat-server/internal/client/db"
)

// CreateMass ...
func (r *Repo) CreateMass(ctx context.Context, userIDs []int64) error {
	countUsers := len(userIDs)
	if countUsers == 0 {
		return nil
	}

	builderInsert := r.sq.Insert(userTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns(idColumn).Suffix("ON CONFLICT DO NOTHING")

	for _, userID := range userIDs {
		builderInsert = builderInsert.Values(userID)
	}

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.CreateMass",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
