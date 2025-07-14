package dto

type BrowseModel[T any] struct {
	RecordsTotal int  `json:"recordsTotal"`
	HasReachMax  bool `json:"hasReachMax"`
	Data         []T  `json:"data"`
}

type Pagination struct {
	SortColumn string `json:"sortColumn"`
	SortOrder  string `json:"sortOrder"`
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
}
