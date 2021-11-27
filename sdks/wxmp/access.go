package wxmp

import (
	"context"
	"time"

	"github.com/szpnygo/mab/internal/errorx"
	"github.com/szpnygo/mab/sdks/wxmp/models"
	"github.com/szpnygo/mab/utils/httpx"
)

const api_access_token = "https://api.weixin.qq.com/cgi-bin/token"

func (helper *WxMpHelper) CheckToken() error {
	_, err := helper.Token()
	if err != nil {
		return err
	}
	return nil
}

func (helper *WxMpHelper) Token() (string, error) {
	tokenKey := "wxmp_token_" + helper.AppID
	if helper.tokenStorage == nil {
		//如果没有储存
		response, err := helper.requestAccessToken()
		if err != nil {
			return "", err
		}
		if !response.IsOK() {
			return "", errorx.New(response.ErrMsg)
		}

		return response.AccessToken, nil
	}

	token, err := helper.tokenStorage.GetToken(context.Background(), tokenKey)
	if err == nil {
		return token, nil
	}
	response, err := helper.requestAccessToken()
	if err != nil {
		return "", err
	}
	if !response.IsOK() {
		return "", errorx.New(response.ErrMsg)
	}

	_ = helper.tokenStorage.SetToken(context.Background(), tokenKey, response.AccessToken, 1*time.Hour)

	return response.AccessToken, nil
}

// requestAccessToken 请求access_token
func (helper *WxMpHelper) requestAccessToken() (token *models.AccessTokenResponse, err error) {
	response, err := httpx.Get(api_access_token).Params(map[string]string{
		"grant_type": "client_credential",
		"appid":      helper.AppID,
		"secret":     helper.AppSecret,
	}).Request()

	if err != nil {
		return nil, err
	}

	return nil, response.JSON(&token)
}
