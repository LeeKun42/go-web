package model

type PaginateInput struct {
	PageSize    int
	CurrentPage int
}

type PaginateOutput struct {
	PageSize    int `json:"page_size"      `
	CurrentPage int `json:"current_page"      `
	Total       int `json:"total"      `
	PageCount   int `json:"page_count"      `
}
