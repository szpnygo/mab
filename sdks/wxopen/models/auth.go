package models

type ComponentAppIDRequest struct {
	ComponentAppID string `json:"component_appid"`
}

type PreAuthCodeResponse struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

type QueryAuthRequest struct {
	ComponentAppID    string `json:"component_appid"`
	AuthorizationCode string `json:"authorization_code"`
}

type RefreshTokenRequest struct {
	ComponentAppID         string `json:"component_appid"`
	AuthorizerAppID        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

type AuthorizerTokenInfo struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"` //接口调用令牌（在授权的公众号/小程序具备 API 权限时，才有此返回值）
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"` //刷新令牌（在授权的公众号具备API权限时，才有此返回值）
}
