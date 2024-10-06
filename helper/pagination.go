package helper

type PaginationParams struct {
	Page         int    `json:"page" query:"page"`
	PageSize     int    `json:"page_size" query:"page_size"`
	SortBy       string `json:"sort_by" query:"sort_by"`
	OrderBy      string `json:"order_by" query:"order_by"`
	FilterBy     string `json:"filter_by" query:"filter_by"`
	FilterValue  string `json:"filter_value" query:"filter_value"`
	TotalRecords int    `json:"total_records"`
	TotalPages   int    `json:"total_pages"`
	Cursor       string `json:"cursor" query:"cursor"`
	HasNextPage  bool   `json:"has_next_page"`
	HasPrevPage  bool   `json:"has_prev_page"`
}

func DefaultPaginationParams() *PaginationParams {
	return &PaginationParams{
		Page:     1,
		PageSize: 10,
		SortBy:   "id",
		OrderBy:  "asc",
	}
}

func (pagination *PaginationParams) ApplyPaginationParams(page int, pageSize int, sortBy string, orderBy string, filterBy string, filterValue string, cursor string) {
	if page > 0 {
		pagination.Page = page
	}

	if pageSize > 0 {
		pagination.PageSize = pageSize
	}

	if sortBy != "" {
		pagination.SortBy = sortBy
	}

	if orderBy != "" {
		pagination.OrderBy = orderBy
	}

	if filterBy != "" {
		pagination.FilterBy = filterBy
	}

	if filterValue != "" {
		pagination.FilterValue = filterValue
	}

	if cursor != "" {
		pagination.Cursor = cursor
	}
}
