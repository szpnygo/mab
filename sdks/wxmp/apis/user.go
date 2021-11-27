package apis

import (
	"github.com/szpnygo/mab/sdks/wxmp/models"
	"github.com/szpnygo/mab/utils/httpx"
)

type UserApi struct {
	token string
}

func NewUserApi(token string) *UserApi {
	return &UserApi{
		token: token,
	}
}

//https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
const api_get_user_info = "https://api.weixin.qq.com/cgi-bin/user/info"

// CreateQRCode 生成用户二维码
func (api *AccountApi) GetUserInfo(openID string) (*models.UserInfo, error) {
	response, err := httpx.Get(api_get_user_info).Params(map[string]string{
		"access_token": api.token,
		"openid":       openID,
	}).Request()
	if err != nil {
		return nil, err
	}
	var result *models.UserInfo
	err = response.JSON(&result)

	return result, err

}
