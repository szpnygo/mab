package wxmp

import "github.com/szpnygo/mab/sdks/token"

type WxMpConfig struct {
	AppID         string
	AppSecret     string
	CallBackToken string
}

type WxMpHelper struct {
	WxMpConfig
	//token保存管理
	tokenStorage token.TokenStorageHelper
}

// NewWxMPHelper 腾讯微信公众号API助手
func NewWxMPHelper(config WxMpConfig) *WxMpHelper {
	return &WxMpHelper{
		WxMpConfig: config,
	}
}

func (h *WxMpHelper) WithTokenStorage(tokenStorage token.TokenStorageHelper) *WxMpHelper {
	h.tokenStorage = tokenStorage
	return h
}
