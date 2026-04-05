package query_builder

import (
	"card/internal/consts"
	sq "github.com/Masterminds/squirrel"
)

func CardQueryBuilder() sq.SelectBuilder {
	return sq.
		Select(
			consts.CardIdColumn,
			consts.CardUserIdColumn,
			consts.CardQuestionColumn,
			consts.CardAnswerColumn,
			consts.CardFileTypeColumn,
			consts.CardFileIdColumn,
			consts.CardCreatedAtColumn,
			consts.CardUpdatedAtColumn,
		).
		From(consts.CardTableName).
		PlaceholderFormat(sq.Dollar)
}
