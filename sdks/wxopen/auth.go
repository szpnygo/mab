package wxopen

import (
	"fmt"

	"github.com/szpnygo/mab/sdks/wxopen/models"
	"github.com/szpnygo/mab/utils/httpx"
)

type AuthType int

const (
	AuthTypeMp      = iota + 1 //只显示公众号
	AuthTypeMiniApp            //只显示小程序
	AUthTypeAll                //全部显示
)

type authApi struct {
	componentAccessToken string // component_access_token
	componentAppID       string // component_appid
}

func newAuthApi(componentAppID, token string) *authApi {
	return &authApi{
		componentAppID:       componentAppID,
		componentAccessToken: token,
	}
}

//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
const api_create_preauthcode = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode"

//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html
const api_query_auth = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth"

// CreatePreAuthCode 预授权码
// 预授权码（pre_auth_code）是第三方平台方实现授权托管的必备信息，每个预授权码有效期为 1800秒。需要先获取令牌才能调用。
func (api *authApi) CreatePreAuthCode() (*models.PreAuthCodeResponse, error) {
	request := &models.ComponentAppIDRequest{
		ComponentAppID: api.componentAppID,
	}
	response, err := httpx.Post(api_create_preauthcode).AddParam("component_access_token", api.componentAccessToken).Struct(request).Request()
	if err != nil {
		return nil, err
	}
	var result *models.PreAuthCodeResponse
	err = response.JSON(&result)

	return result, err
}

// CreatePreAuthURL 构建PC端授权链接
func (api *authApi) CreatePCPreAuthURL(authType AuthType, code string, callback string) string {
	return fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d",
		api.componentAppID, code, callback, authType)
}

// QueryAuth s使用授权码获取授权信息
// 当用户在第三方平台授权页中完成授权流程后，第三方平台开发者可以在回调 URI 中通过 URL 参数获取授权码。
// 使用以下接口可以换取公众号/小程序的授权信息。建议保存授权信息中的刷新令牌（authorizer_refresh_token）
func (api *authApi) QueryAuth(code string) (*models.AuthorizationInfo, error) {
	request := &models.QueryAuthRequest{
		ComponentAppID:    api.componentAppID,
		AuthorizationCode: code,
	}
	response, err := httpx.Post(api_query_auth).AddParam("component_access_token", api.componentAccessToken).Struct(request).Request()
	if err != nil {
		return nil, err
	}
	var result *models.AuthorizationInfo
	err = response.JSON(&result)

	return result, err
}
