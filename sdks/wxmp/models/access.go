package models

// AccessTokenResponse 公众号token返回
type AccessTokenResponse struct {
	BaseResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
