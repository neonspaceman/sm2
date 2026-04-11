package query

import (
	"card/internal/consts"
	domain_card "card/internal/domain/card"
	"card/internal/query_builder"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GetCardsByUserIdQuery struct {
	UserId string `validate:"required,uuid"`
	Limit  uint64 `validate:"required"`
	After  string `validate:"omitempty,uuid"`
}

type GetCardByUserIdHandler struct {
	conn     *pgxpool.Pool
	validate *validator.Validate
}

func NewGetCardByUserIdHandler(
	conn *pgxpool.Pool,
	validate *validator.Validate,
) *GetCardByUserIdHandler {
	return &GetCardByUserIdHandler{
		conn:     conn,
		validate: validate,
	}
}

func (h *GetCardByUserIdHandler) Handle(ctx context.Context, cmd GetCardsByUserIdQuery) ([]*domain_card.Card, error) {
	err := h.validate.Struct(&cmd)

	if err != nil {
		return nil, err
	}

	b := query_builder.CardQueryBuilder().
		Where(sq.Eq{consts.CardUserIdColumn: cmd.UserId}).
		OrderBy(fmt.Sprintf("%s DESC", consts.CardIdColumn)).
		Limit(cmd.Limit)

	if cmd.After != "" {
		b = b.Where(sq.Lt{consts.CardIdColumn: cmd.After})
	}

	sql, args := b.MustSql()

	rows, err := h.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	cards, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain_card.Card])

	if err != nil {
		return nil, fmt.Errorf("collect one: %w", err)
	}

	return cards, nil
}
