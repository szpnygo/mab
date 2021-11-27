package ffapp

import (
	"github.com/szpnygo/mab/internal/errorx"
	"github.com/szpnygo/mab/utils/httpx"
)

const (
	// 消息推送
	API_FF_APP_SEND_MESSAGE = "https://open.feishu.cn/open-apis/im/v1/messages"
)

// BotPush 机器人消息推送
func (m *FFApp) BotPush(receive string, receiveType string, data string, msgType string) (*MessageData, error) {
	request := SendMsgRequest{
		ReceiveID: receive,
		Content:   data,
		MsgType:   msgType,
	}
	var response MessageResponse

	err := httpx.Post(API_FF_APP_SEND_MESSAGE).AddParam("receive_id_type", receiveType).AddHeader("Authorization", "Bearer "+m.token).
		Struct(request).JSON(&response)

	if err != nil {
		return nil, err
	}

	if response.Code != 0 {
		return nil, errorx.Errorf("%d:%s", response.Code, response.Msg)
	}

	return &response.Data, nil
}
