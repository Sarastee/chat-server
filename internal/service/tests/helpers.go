package tests

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/chat-server/internal/service/mocks"
	dbMocks "github.com/sarastee/platform_common/pkg/db/mocks"
	"github.com/sarastee/platform_common/pkg/db/pg"
)

func txFakerAndCtxWithSetup(ctx context.Context, t *testing.T, successTx bool) (*mocks.TxFaker, context.Context) {
	txFaker := mocks.NewTxFaker(t)
	ctxWithTx := pg.MakeContextTx(ctx, txFaker)

	if successTx {
		txFaker.On("Commit", ctxWithTx).Return(nil).Once()
	} else {
		txFaker.On("Rollback", ctxWithTx).Return(nil).Once()
	}

	return txFaker, ctxWithTx
}

func transactorWithSetup(ctx context.Context, t *testing.T, faker *mocks.TxFaker) *dbMocks.Transactor {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	transactorMock := dbMocks.NewTransactor(t)
	transactorMock.On("BeginTx", ctx, txOpts).Return(faker, nil).Once()

	return transactorMock
}
