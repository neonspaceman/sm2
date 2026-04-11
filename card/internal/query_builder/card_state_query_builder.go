package query_builder

import (
	"card/internal/consts"
	sq "github.com/Masterminds/squirrel"
)

func CardStateQueryBuilder() sq.SelectBuilder {
	return sq.
		Select(
			consts.CardStateIdColumn,
			consts.CardStateStateColumn,
			consts.CardStateStepColumn,
			consts.CardStateEasyColumn,
			consts.CardStateDueColumn,
			consts.CardStateCurrentIntervalInDaysColumn,
			consts.CardStateCreatedAtColumn,
			consts.CardStateUpdatedAtColumn,
		).
		From(consts.CardStateTableName).
		PlaceholderFormat(sq.Dollar)
}
