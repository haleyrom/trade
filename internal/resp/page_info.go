package resp

// PageInfoResp 分页信息参数
type PageInfoResp struct {
	PageSize int `json:"page_size"`
	Count    int `json:"count"`
}

// PageResp 分页参数
type PageResp struct {
	Items interface{}  `json:"items"`
	Page  PageInfoResp `json:"page"`
}
