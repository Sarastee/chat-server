package chat

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/chat-server/internal/client/db"
)

// IsUserInChat ...
func (r *Repo) IsUserInChat(ctx context.Context, chatID int64, userID int64) (bool, error) {
	builderIsUserInChat := r.sq.Select(chatIDColumn).PlaceholderFormat(squirrel.Dollar).Where(squirrel.Eq{
		userIDColumn: userID,
		chatIDColumn: chatID,
	})

	query, args, err := builderIsUserInChat.ToSql()
	if err != nil {
		return false, err
	}

	q := db.Query{
		Name:     "chat_repository.IsUserInChat",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	_, err = pgx.CollectOneRow(rows, pgx.RowTo[bool])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
