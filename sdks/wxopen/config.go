package wxopen

import "github.com/szpnygo/mab/sdks/token"

type WxOpenConfig struct {
	ComponentAppID     string
	ComponentAppSecret string
	CallbackToken      string
	EncodingAESKey     string
}

type WxOpenHelper struct {
	WxOpenConfig
	//token保存管理
	tokenStorage token.TokenStorageHelper
}

// NewWxMPHelper 微信开放平台API助手
func NewWxMPHelper(config WxOpenConfig) *WxOpenHelper {
	return &WxOpenHelper{
		WxOpenConfig: config,
	}
}

func (h *WxOpenHelper) WithTokenStorage(tokenStorage token.TokenStorageHelper) *WxOpenHelper {
	h.tokenStorage = tokenStorage
	return h
}
