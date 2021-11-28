package wxopen

import (
	"context"
	"time"

	"github.com/szpnygo/mab/internal/errorx"
	"github.com/szpnygo/mab/sdks/wxopen/models"
	"github.com/szpnygo/mab/utils/httpx"
)

const api_component_access_token = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"

func (helper *WxOpenHelper) CheckToken() error {
	_, err := helper.Token()
	if err != nil {
		return err
	}
	return nil
}

func (helper *WxOpenHelper) Token() (string, error) {
	tokenKey := "wxopen_component_token_" + helper.ComponentAppID
	if helper.tokenStorage == nil {
		//如果没有储存
		response, err := helper.requestAccessToken()
		if err != nil {
			return "", err
		}
		if !response.IsOK() {
			return "", errorx.New(response.ErrMsg)
		}

		return response.ComponentAccessToken, nil
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

	_ = helper.tokenStorage.SetToken(context.Background(), tokenKey, response.ComponentAccessToken, 1*time.Hour)

	return response.ComponentAccessToken, nil
}

// requestAccessToken 请求access_token
func (helper *WxOpenHelper) requestAccessToken() (token *models.ComponentTokenResponse, err error) {
	request := models.ComponentTokenRequest{
		ComponentAppID:        helper.ComponentAppID,
		ComponentAppSecret:    helper.ComponentAppSecret,
		ComponentVerifyTicket: helper.getTicket(),
	}
	err = httpx.Post(api_component_access_token).Struct(&request).JSON(&token)

	return
}
