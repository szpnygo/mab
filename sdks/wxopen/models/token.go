package models

type ComponentTokenRequest struct {
	ComponentAppID        string `json:"component_appid"`
	ComponentAppSecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

type ComponentTokenResponse struct {
	BaseResponse
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int    `json:"expires_in"`
}
