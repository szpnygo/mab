package models

type BaseResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (base *BaseResponse) IsOK() bool {
	return base.ErrCode == 0
}
