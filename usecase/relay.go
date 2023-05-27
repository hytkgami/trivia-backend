package usecase

type CursorParams struct {
	OrderBy string
	First   *int
	After   *string
	Last    *int
	Before  *string
}

type CursorPagination struct {
	Cursor         *string
	Limit          int
	OrderBy        string
	OrderDirection string
}

const (
	OrderDirectionASC  = "ASC"
	OrderDirectionDESC = "DESC"
)

func (p *CursorParams) ToCursorPagination() *CursorPagination {
	var cursor *string
	limit := p.First
	orderDirection := OrderDirectionASC
	if p.Last != nil {
		limit = p.Last
		orderDirection = OrderDirectionDESC
	}
	if p.After != nil {
		cursor = p.After
	} else if p.Before != nil {
		cursor = p.Before
	}
	return &CursorPagination{
		Cursor:         cursor,
		Limit:          *limit,
		OrderBy:        p.OrderBy,
		OrderDirection: orderDirection,
	}
}

type PageInfo struct {
	HasNextPage bool
	Cursor      string
}
