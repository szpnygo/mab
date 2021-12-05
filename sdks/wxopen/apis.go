package wxopen

import (
	"log"

	"github.com/szpnygo/mab/internal/errorx"
	"github.com/szpnygo/mab/sdks/wxopen/apis"
)

// NewAuthApi 授权相关接口
func (helper *WxOpenHelper) NewAuthApi() *apis.AuthApi {
	token, err := helper.Token()
	if err != nil {
		errorx.Log(err)
		log.Fatalf(err.Error())
	}

	return apis.NewAuthApi(helper.ComponentAppID, token)
}
