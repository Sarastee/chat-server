package message

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sarastee/chat-server/internal/client/db"
	serviceModel "github.com/sarastee/chat-server/internal/model"
	"github.com/sarastee/chat-server/internal/repository"
	"github.com/sarastee/chat-server/internal/repository/message/convert"
)

// Create ...
func (r *Repo) Create(ctx context.Context, message serviceModel.Message) error {
	repoMessage := convert.ToMessageFromServiceMessage(&message)

	builderInsert := r.sq.Insert(messageTable).PlaceholderFormat(squirrel.Dollar).
		Columns(fromUserIDColumn, chatIDColumn, textColumn, sentAtColumn).
		Values(repoMessage.FromUserID, repoMessage.ToChatID, repoMessage.Text, repoMessage.SendTime)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "message_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args)
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
	}

	return nil
}
