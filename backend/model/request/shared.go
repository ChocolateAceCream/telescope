/*
* @fileName shared.go
* @author Di Sheng
* @date 2022/11/28 15:41:44
* @description shared request model
 */

package request

type CursorListParam struct {
	Pagination
	CursorId uint `json:"cursor_id" form:"cursor_id"`
	// Title    string `json:"title" form:"title"`
	// Author   string `json:"author" form:"author"`
	Desc bool `json:"desc" form:"desc"` // order by desc (by default)
}

type Pagination struct {
	PageNumber int `json:"page_number" form:"page_number"`
	PageSize   int `json:"page_size" form:"page_size"`
}

type FindById struct {
	ID int `json:"id" form:"id"`
}

type FindByIds struct {
	ID []int `json:"id" form:"id"`
}
