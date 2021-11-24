package ffapp

import (
	"fmt"

	"github.com/szpnygo/mab/utils/httpx"
)

const (
	// 获取用户信息
	API_FF_APP_AUTH_USER_INFO = "https://open.feishu.cn/open-apis/authen/v1/access_token"
)

// AuthUserInfo 授权的用户信息
func (m *FFApp) AuthUserInfo(code string) (*UserInfo, error) {
	request := AuthUserInfoRequest{
		Code:      code,
		GrantType: "authorization_code",
	}
	var response AuthUserInfoResponse

	err := httpx.Get(API_FF_APP_AUTH_USER_INFO).AddHeader("Authorization", "Bearer "+m.token).
		Struct(&request).JSON(&response)

	if err != nil {
		return nil, err
	}

	if response.Code != 0 {
		return nil, fmt.Errorf("%d:%s", response.Code, response.Msg)
	}

	return &response.Data, nil
}
