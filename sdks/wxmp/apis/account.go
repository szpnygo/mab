package apis

import (
	"github.com/szpnygo/mab/sdks/wxmp/models"
	"github.com/szpnygo/mab/utils/httpx"
)

type AccountApi struct {
	token string
}

func NewAccountApi(token string) *AccountApi {
	return &AccountApi{
		token: token,
	}
}

//https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
const api_create_qr_code = "https://api.weixin.qq.com/cgi-bin/qrcode/create"

// CreateQRCode 生成用户二维码
func (api *AccountApi) CreateQRCode(ticket string) (*models.QRResult, error) {
	response, err := httpx.Get(api_create_qr_code).AddParam("access_token", api.token).Request()
	if err != nil {
		return nil, err
	}
	var result *models.QRResult
	err = response.JSON(&result)

	return result, err

}
