package resp

import "github.com/oceanho/gcontrib/webapi"

// NewResp returns a webapi.Resp Object.
func NewResp(succ bool, code int32, message string) *webapi.Resp {
	return &webapi.Resp{
		Succ:    succ,
		Code:    code,
		Message: message,
	}
}

// NewErr returns a error webapi.Resp Object.
func NewErr(code int32, err error) *webapi.Resp {
	return NewResp(false, code, err.Error())
}

// NewOK returns a successfully webapi.Resp Object.
func NewOK(code int32, message string) *webapi.Resp {
	return NewResp(true, code, message)
}

// NewOKDef returns a default success webapi.Resp Object.
func NewOKDef() *webapi.Resp {
	return NewResp(true, 0, "")
}

// NewResult returns a default success webapi.RespResult Object.
func NewResult(result interface{}) *webapi.RespResult {
	return &webapi.RespResult{
		Result: result,
		Resp:   *NewOKDef(),
	}
}

// NewRespPager returns a success pager webapi.RespResult Object.
func NewRespPager(pager webapi.ReqPager, total int64, pageCount int32, data interface{}) *webapi.RespResult {
	return NewResult(webapi.PagerData{
		ReqPager:  pager,
		Total:     total,
		PageCount: pageCount,
		Data:      data,
	})
}
