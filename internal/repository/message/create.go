package message

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
	serviceModel "github.com/sarastee/chat-server/internal/model"
	"github.com/sarastee/chat-server/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

// Create ...
func (r *Repo) Create(ctx context.Context, message serviceModel.Message) error {
	builderInsert := r.sq.Insert(messageTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns(fromUserIDColumn, chatIDColumn, textColumn, sentAtColumn).
		Values(message.FromUserID, message.ToChatID, message.Text, message.SendTime)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "message_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch {
			case pgErr.ConstraintName == messageChatIDFKConstraint:
				return repository.ErrChatNotFound
			case pgErr.ConstraintName == messageFromUserIDFKConstraint:
				return repository.ErrUserNotFound
			}
		}
		return err
	}

	return nil
}
