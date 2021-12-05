package wxopen

import (
	"log"

	"github.com/szpnygo/mab/internal/errorx"
)

// NewAuthApi 授权相关接口
func (helper *WxOpenHelper) NewAuthApi() *authApi {
	token, err := helper.Token()
	if err != nil {
		errorx.Log(err)
		log.Fatalf(err.Error())
	}

	return newAuthApi(helper.ComponentAppID, token)
}
