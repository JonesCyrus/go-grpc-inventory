package models

type Pagination struct {
	Search string `json:"search"`
	Page   int    `json:"page"`
	Size   int    `json:"size"`
}
