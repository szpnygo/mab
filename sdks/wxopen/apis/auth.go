package apis

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

type AuthApi struct {
	componentAccessToken string // component_access_token
	componentAppID       string // component_appid
}

func NewAuthApi(componentAppID, token string) *AuthApi {
	return &AuthApi{
		componentAppID:       componentAppID,
		componentAccessToken: token,
	}
}

const (
	api_create_preauthcode  = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode"
	api_query_auth          = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth"
	api_authorizer_token    = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token"
	api_get_authorizer_info = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info"
)

// CreatePreAuthCode 预授权码
// 预授权码（pre_auth_code）是第三方平台方实现授权托管的必备信息，每个预授权码有效期为 1800秒。需要先获取令牌才能调用。
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
func (api *AuthApi) CreatePreAuthCode() (*models.PreAuthCodeResponse, error) {
	request := &models.ComponentAppIDRequest{
		ComponentAppID: api.componentAppID,
	}
	response, err := httpx.Post(api_create_preauthcode).AddParam("component_access_token", api.componentAccessToken).Struct(request).Request()
	if err != nil {
		return nil, err
	}
	var result *models.PreAuthCodeResponse
	err = response.JSON(&result)
	if err != nil {
		return result, nil
	}

	return nil, err
}

// CreatePreAuthURL 构建PC端授权链接
func (api *AuthApi) CreatePCPreAuthURL(authType AuthType, code string, callback string) string {
	return fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d",
		api.componentAppID, code, callback, authType)
}

// QueryAuth 使用授权码获取授权信息
// 当用户在第三方平台授权页中完成授权流程后，第三方平台开发者可以在回调 URI 中通过 URL 参数获取授权码。
// 使用以下接口可以换取公众号/小程序的授权信息。建议保存授权信息中的刷新令牌（authorizer_refresh_token）
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html
func (api *AuthApi) QueryAuth(code string) (*models.AuthorizationInfo, error) {
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
	if err != nil {
		return result, nil
	}

	return nil, err
}

// RefreshToken 获取/刷新接口调用令牌
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html
func (api *AuthApi) RefreshToken(authorizerAppID string, authorizerRefreshToken string) (*models.AuthorizerTokenInfo, error) {
	request := &models.RefreshTokenRequest{
		ComponentAppID:         api.componentAppID,
		AuthorizerAppID:        authorizerAppID,
		AuthorizerRefreshToken: authorizerRefreshToken,
	}
	response, err := httpx.Post(api_authorizer_token).AddParam("component_access_token", api.componentAccessToken).Struct(request).Request()
	if err != nil {
		return nil, err
	}
	var result *models.AuthorizerTokenInfo
	err = response.JSON(&result)
	if err != nil {
		return result, nil
	}
	return nil, err
}

// GetAuthorizerInfo 获取授权方的帐号基本信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_get_authorizer_info.html
func (api *AuthApi) GetAuthorizerInfo(authorizerAppID string) (*models.AuthInfoResult, error) {
	request := &models.GetAuthInfoRequest{
		ComponentAppID:  api.componentAppID,
		AuthorizerAppID: authorizerAppID,
	}

	response, err := httpx.Post(api_get_authorizer_info).AddParam("component_access_token", api.componentAccessToken).Struct(request).Request()
	if err != nil {
		return nil, err
	}
	var result *models.AuthInfoResult
	err = response.JSON(&result)
	if err != nil {
		return result, nil
	}

	return nil, err
}
