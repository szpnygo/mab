package ffapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/gogf/gf/util/gconv"

	"github.com/szpnygo/mab/utils/configx"
)

const (
	API_FF_APP_ACCESS_TOKEN   = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal"
	API_FF_APP_AUTH_USER_INFO = "https://open.feishu.cn/open-apis/authen/v1/access_token"
	API_FF_APP_SEND_MESSAGE   = "https://open.feishu.cn/open-apis/im/v1/messages" // 消息推送
	API_FF_APP_CHATS          = "https://open.feishu.cn/open-apis/im/v1/chats"    // 获取用户或机器人所在的群列表
)

type FFApp struct {
	Config *configx.FFAppConf
	token  string
}

func NewFFAppHelper(config *configx.FFAppConf) *FFApp {
	return &FFApp{
		Config: config,
	}
}

func (m *FFApp) SetAccessToken(token string) {
	m.token = token
}

func (m *FFApp) GetAccessToken() (string, error) {
	if len(m.token) > 0 {
		return m.token, nil
	}

	accessToken, err := m.RequestAccessToken()
	if err != nil {
		return "", err
	}

	return accessToken.AppAccessToken, nil
}

func (m *FFApp) AccessToken() {
	_, _ = m.RequestAccessToken()
}

func (m *FFApp) RequestAccessToken() (*AccessTokenResponse, error) {
	request := AccessTokenRequest{
		AppID:     m.Config.AppID,
		AppSecret: m.Config.AppSecret,
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := PostRaw(API_FF_APP_ACCESS_TOKEN, requestData, nil)
	if err != nil {
		return nil, err
	}

	var access AccessTokenResponse
	err = json.Unmarshal(body, &access)
	if err != nil {
		return nil, err
	}

	if access.Code != 0 {
		return nil, fmt.Errorf("%d:%s", access.Code, access.Msg)
	}

	m.token = access.AppAccessToken

	return &access, nil
}

func (m *FFApp) AuthUserInfo(code string) (*UserInfo, error) {
	request := AuthUserInfoRequest{
		Code:      code,
		GrantType: "authorization_code",
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := PostRaw(API_FF_APP_AUTH_USER_INFO, requestData, map[string]string{
		"Authorization": "Bearer " + m.token,
	})
	if err != nil {
		return nil, err
	}

	var access AuthUserInfoResponse
	err = json.Unmarshal(body, &access)
	if err != nil {
		return nil, err
	}

	if access.Code != 0 {
		return nil, fmt.Errorf("%d:%s", access.Code, access.Msg)
	}

	return &access.Data, nil
}

// BotPush 机器人消息推送
func (m *FFApp) BotPush(receive string, receiveType string, data string, msgType string) (*MessageData, error) {
	request := SendMsgRequest{
		ReceiveID: receive,
		Content:   data,
		MsgType:   msgType,
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := PostRaw(API_FF_APP_SEND_MESSAGE+"?receive_id_type="+receiveType, requestData, map[string]string{
		"Authorization": "Bearer " + m.token,
	})
	if err != nil {
		return nil, err
	}

	var msg MessageResponse
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return nil, err
	}

	if msg.Code != 0 {
		return nil, fmt.Errorf("%d:%s", msg.Code, msg.Msg)
	}

	return &msg.Data, nil
}

// Chats 获取机器人或者用户所在群组
func (m *FFApp) Chats(pageSize int, pageToken ...string) (*Chats, error) {
	values := url.Values{}
	values.Set("page_size", gconv.String(pageSize))
	if len(pageToken) > 0 {
		values.Set("page_token", gconv.String(pageToken[0]))
	}
	data, err := Get(API_FF_APP_CHATS+"?"+values.Encode(), map[string]string{
		"Authorization": "Bearer " + m.token,
		"Content-Type":  "application/json; charset=utf-8",
	})
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}
	respData := &Chats{}
	if err := json.Unmarshal(resp.Data, respData); err != nil {
		return nil, err
	}
	return respData, nil
}
