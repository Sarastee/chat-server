package chat

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/sarastee/chat-server/internal/client/db"
)

// Delete ...
func (r *Repo) Delete(ctx context.Context, chatID int64) error {
	builderDelete := r.sq.Delete(chatTable).PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: chatID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
