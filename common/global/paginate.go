package global

// PageData 公共分页数据返回
type PageData struct {
	//  总条数
	Total int64 `json:"total"`
	// 是否分页
	IsPage *bool `json:"isPage"`
	// 当前页
	Page int `json:"page"`
	// 每页条数
	PageSize int `json:"pageSize"`
	// 数据列表
	List interface{} `json:"list"`
}
