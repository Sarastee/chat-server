package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/sarastee/chat-server/internal/client/db"
)

// CreateMass ...
func (r *Repo) CreateMass(ctx context.Context, userIDs []int64) error {
	countUsers := len(userIDs)
	if countUsers == 0 {
		return nil
	}

	var strUserIds = make([]string, 0, countUsers)
	for _, userID := range userIDs {
		strUserIds = append(strUserIds, fmt.Sprintf("(%d)", userID))
	}

	values := strings.Join(strUserIds, ",")
	builderInsertMass := r.sq.Insert(userTable).PlaceholderFormat(squirrel.Dollar).Into(idColumn).
		Values(values)

	query, args, err := builderInsertMass.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "user_repository.CreateMass",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args)
	if err != nil {
		return err
	}

	return nil
}
