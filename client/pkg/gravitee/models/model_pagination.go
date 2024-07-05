package models

type Pagination struct {
	Page           int   `json:"page"`
	PerPage        int   `json:"perPage"`
	PageCount      int   `json:"pageCount"`
	PageItemsCount int   `json:"pageItemsCount"`
	TotalCount     int64 `json:"totalCount"`
}
