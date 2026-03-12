package shared

type Pagination struct {
	Page     int32
	PageSize int32
}

func NewPagination(page, pageSize int32) Pagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	return Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

func (p Pagination) Offset() int32 {
	return (p.Page - 1) * p.PageSize
}

func (p Pagination) Limit() int32 {
	return p.PageSize
}

type PaginatedResult[T any] struct {
	Items      []T
	TotalCount int64
	Page       int32
	PageSize   int32
	TotalPages int32
}

func NewPaginatedResult[T any](items []T, totalCount int64, pagination Pagination) PaginatedResult[T] {
	totalPages := int32(totalCount) / pagination.PageSize
	if int32(totalCount)%pagination.PageSize > 0 {
		totalPages++
	}

	return PaginatedResult[T]{
		Items:      items,
		TotalCount: totalCount,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: totalPages,
	}
}
