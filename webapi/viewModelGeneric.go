package webapi

// Resp represents a generic Web API response Object
type Resp struct {
	Succ    bool   `json:"succ"`
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// ReqPager represents a generic Web API pager request Object
type ReqPager struct {
	PageNumber int32 `json:"page_no"`
	PageSize   int32 `json:"page_size"`
}

// PagerData represents a generic Web API pager data result Object
type PagerData struct {
	ReqPager
	PageCount int32       `json:"page_count"`
	Total     int64       `json:"total"`
	Data      interface{} `json:"data"`
}

// RespResult represents a generic result Web API response Object
type RespResult struct {
	Resp
	Result interface{} `json:"result"`
}
