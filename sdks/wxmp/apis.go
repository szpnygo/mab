package wxmp

import (
	"log"

	"github.com/szpnygo/mab/internal/errorx"
	"github.com/szpnygo/mab/sdks/wxmp/apis"
)

func (helper *WxMpHelper) AccountApi() *apis.AccountApi {
	token, err := helper.Token()
	if err != nil {
		errorx.Log(err)
		log.Fatalf(err.Error())
	}

	return apis.NewAccountApi(token)
}
