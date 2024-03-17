package chat

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/platform_common/pkg/db"
)

// IsUserInChat ...
func (r *Repo) IsUserInChat(ctx context.Context, chatID int64, userID int64) (bool, error) {
	// `SELECT TRUE FROM %s WHERE %s = @%s AND %s = @%s`

	builderSelect := r.sq.Select("TRUE").
		PlaceholderFormat(squirrel.Dollar).
		From(chatUserTable).
		Where(squirrel.Eq{
			chatIDColumn: chatID,
			userIDColumn: userID,
		})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return false, err
	}

	q := db.Query{
		Name:     "chat_repository.IsUserInChat",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
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
