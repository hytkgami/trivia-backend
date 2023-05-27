package repository

import (
	"fmt"

	"github.com/hytkgami/trivia-backend/usecase"
	"github.com/jmoiron/sqlx"
)

func generateCursorQuery(p *usecase.CursorPagination) (string, []any, error) {
	switch p.OrderDirection {
	case usecase.OrderDirectionASC:
		if p.Cursor != nil {
			query := fmt.Sprintf(" AND :order_by > :cursor ORDER BY :order_by %s LIMIT :limit", p.OrderDirection)
			return sqlx.Named(query, map[string]any{
				"order_by": p.OrderBy,
				"cursor":   *p.Cursor,
				"limit":    p.Limit + 1, // for hasNextPage
			})
		}
		query := fmt.Sprintf(" ORDER BY :order_by %s LIMIT :limit", p.OrderDirection)
		return sqlx.Named(query, map[string]any{
			"order_by": p.OrderBy,
			"limit":    p.Limit + 1, // for hasNextPage
		})
	case usecase.OrderDirectionDESC:
		if p.Cursor != nil {
			query := fmt.Sprintf(" AND :order_by > :cursor ORDER BY :order_by %s LIMIT :limit", p.OrderDirection)
			return sqlx.Named(query, map[string]any{
				"order_by": p.OrderBy,
				"cursor":   *p.Cursor,
				"limit":    p.Limit + 1, // for hasNextPage
			})
		}
		query := fmt.Sprintf(" ORDER BY :order_by %s LIMIT :limit", p.OrderDirection)
		return sqlx.Named(query, map[string]any{
			"order_by": p.OrderBy,
			"limit":    p.Limit + 1, // for hasNextPage
		})
	}
	return "", nil, fmt.Errorf("invalid cursor params: first or last must be set")
}
