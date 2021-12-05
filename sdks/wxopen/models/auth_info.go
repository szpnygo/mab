package models

type AuthorizationInfo struct {
	BaseResponse
	AuthorizerAppID        string     `json:"authorizer_appid"`        //授权方 appid
	AuthorizerAccessToken  string     `json:"authorizer_access_token"` //接口调用令牌（在授权的公众号/小程序具备 API 权限时，才有此返回值）
	ExpiresIn              int        `json:"expires_in"`
	AuthorizerRefreshToken string     `json:"authorizer_refresh_token"` //刷新令牌（在授权的公众号具备API权限时，才有此返回值）
	FuncInfo               []FuncInfo `json:"func_info"`                //授权给开发者的权限集列表
}

type FuncInfo struct {
	FuncscopeCategory FuncscopeCategory `json:"funcscope_category"`
}

type FuncscopeCategory struct {
	ID int `json:"id"`
}
