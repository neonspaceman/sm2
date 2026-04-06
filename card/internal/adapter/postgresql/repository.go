package postgresql

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
	trm  *trmpgx.CtxGetter
}

func (r *Repository) conn(ctx context.Context) trmpgx.Tr {
	return r.trm.DefaultTrOrDB(ctx, r.pool)
}
