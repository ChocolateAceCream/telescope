package request

import "strings"

type ProjectListRequest struct {
	Params ProjectListRequestParam `json:"params" form:"params"`
}

type ProjectListRequestParam struct {
	PageNumber int    `json:"page_number" form:"page_number"`
	PageSize   int    `json:"page_size" form:"page_size"`
	OrderBy    string `json:"order_by" form:"order_by"`
	SortBy     string `json:"sort_by" form:"sort_by"`
}

func (r *ProjectListRequestParam) ApplyDefaultsAndValidate() error {
	// Default values
	if r.PageNumber <= 0 {
		r.PageNumber = 1
	}
	if r.PageSize <= 0 || r.PageSize > 100 {
		r.PageSize = 15
	}

	// Whitelist allowed order fields
	allowedOrderFields := map[string]bool{
		"updated_at":   true,
		"created_at":   true,
		"project_name": true,
	}
	if !allowedOrderFields[r.OrderBy] {
		r.OrderBy = "updated_at"
	}

	// Whitelist allowed sort direction
	r.SortBy = strings.ToLower(r.SortBy)
	if r.SortBy != "asc" && r.SortBy != "desc" {
		r.SortBy = "desc"
	}

	return nil
}
