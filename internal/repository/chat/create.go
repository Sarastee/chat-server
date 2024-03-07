package chat

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sarastee/chat-server/internal/client/db"
)

// Create chat in db
func (r *Repo) Create(ctx context.Context) (int64, error) {
	// `insert into %s default values returning id`

	buildInsert := r.sq.Insert(chatTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns(idColumn).
		Values(squirrel.Expr("DEFAULT")).
		Suffix("RETURNING id")

	query, _, err := buildInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	chatID, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
