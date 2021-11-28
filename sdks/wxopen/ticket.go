package wxopen

import (
	"context"

	"github.com/szpnygo/mab/internal/errorx"
)

var appTicket string = ""

func (helper *WxOpenHelper) SaveTicket(data []byte, msgSignature string, timestamp string, nonce string) {
	message, err := helper.MessageCallback(data, msgSignature, timestamp, nonce)
	if err != nil {
		return
	}
	if message.InfoType != "component_verify_ticket" {
		return
	}

	key := "wxopen_ticket_" + helper.ComponentAppID
	if helper.tokenStorage != nil {
		err := helper.tokenStorage.SetToken(context.Background(), key, message.ComponentVerifyTicket, 7200)
		if err != nil {
			errorx.Log(err)
		}
	} else {
		appTicket = message.ComponentVerifyTicket
	}
}

func (helper *WxOpenHelper) getTicket() string {
	key := "wxopen_ticket_" + helper.ComponentAppID
	if helper.tokenStorage != nil {
		ticket, err := helper.tokenStorage.GetToken(context.Background(), key)
		if err != nil {
			errorx.Log(err)
			return appTicket
		}
		return ticket
	}

	return appTicket
}
