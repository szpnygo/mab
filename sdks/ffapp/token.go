package ffapp

import (
	"github.com/szpnygo/mab/internal/errorx"
	"github.com/szpnygo/mab/utils/httpx"
)

const (
	// 获取access_token
	API_FF_APP_ACCESS_TOKEN = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal"
)

// RequestAccessToken 请求AccessToken
func (m *FFApp) RequestAccessToken() (*AccessTokenResponse, error) {
	request := AccessTokenRequest{
		AppID:     m.Config.AppID,
		AppSecret: m.Config.AppSecret,
	}
	var response AccessTokenResponse

	err := httpx.Post(API_FF_APP_ACCESS_TOKEN).Struct(&request).JSON(&response)
	if err != nil {
		return nil, err
	}

	if response.Code != 0 {
		return nil, errorx.Errorf("%d:%s", response.Code, response.Msg)
	}

	m.token = response.AppAccessToken

	return &response, nil
}

func (m *FFApp) SetAccessToken(token string) {
	m.token = token
}

func (m *FFApp) GetAccessToken() (string, error) {
	if len(m.token) > 0 {
		return m.token, nil
	}

	accessToken, err := m.RequestAccessToken()
	if err != nil {
		return "", err
	}

	return accessToken.AppAccessToken, nil
}

func (m *FFApp) AccessToken() {
	_, _ = m.RequestAccessToken()
}
