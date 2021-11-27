package ffapp

import "encoding/json"

type AccessTokenRequest struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type AccessTokenResponse struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}

type AuthUserInfoResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data UserInfo `json:"data"`
}
type UserInfo struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Name             string `json:"name"`
	EnName           string `json:"en_name"`
	AvatarURL        string `json:"avatar_url"`
	AvatarThumb      string `json:"avatar_thumb"`
	AvatarMiddle     string `json:"avatar_middle"`
	AvatarBig        string `json:"avatar_big"`
	OpenID           string `json:"open_id"`
	UnionID          string `json:"union_id"`
	Email            string `json:"email"`
	UserID           string `json:"user_id"`
	Mobile           string `json:"mobile"`
	TenantKey        string `json:"tenant_key"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
}

type AuthUserInfoRequest struct {
	GrantType string `json:"grant_type"`
	Code      string `json:"code"`
}

const (
	SEND_MSG_RECEIVE_ID_TYPE_OPEN_ID  = "open_id"
	SEND_MSG_RECEIVE_ID_TYPE_USER_ID  = "user_id"
	SEND_MSG_RECEIVE_ID_TYPE_UNION_ID = "union_id"
	SEND_MSG_RECEIVE_ID_TYPE_EMAIL    = "email"
	SEND_MSG_RECEIVE_ID_TYPE_CHAT_ID  = "chat_id"
)

const (
	MESSAGE_TYPE_TEXT        = "text"
	MESSAGE_TYPE_POST        = "post"
	MESSAGE_TYPE_IMAGE       = "image"
	MESSAGE_TYPE_INTERACTIVE = "interactive"
	MESSAGE_TYPE_SHARE_CHAT  = "share_chat"
	MESSAGE_TYPE_SHARE_USER  = "share_user"
	MESSAGE_TYPE_AUDIO       = "audio"
	MESSAGE_TYPE_MEDIA       = "media"
	MESSAGE_TYPE_FILE        = "file"
	MESSAGE_TYPE_STICKER     = "sticker"
)

type SendMsgRequest struct {
	ReceiveID string `json:"receive_id"`
	Content   string `json:"content"`
	MsgType   string `json:"msg_type"`
}

type MessageResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data MessageData `json:"data"`
}
type Sender struct {
	ID         string `json:"id"`
	IDType     string `json:"id_type"`
	SenderType string `json:"sender_type"`
	TenantKey  string `json:"tenant_key"`
}
type Body struct {
	Content string `json:"content"`
}
type Mentions struct {
	Key       string `json:"key"`
	ID        string `json:"id"`
	IDType    string `json:"id_type"`
	Name      string `json:"name"`
	TenantKey string `json:"tenant_key"`
}
type MessageData struct {
	MessageID      string     `json:"message_id"`
	RootID         string     `json:"root_id"`
	ParentID       string     `json:"parent_id"`
	MsgType        string     `json:"msg_type"`
	CreateTime     string     `json:"create_time"`
	UpdateTime     string     `json:"update_time"`
	Deleted        bool       `json:"deleted"`
	Updated        bool       `json:"updated"`
	ChatID         string     `json:"chat_id"`
	Sender         Sender     `json:"sender"`
	Body           Body       `json:"body"`
	Mentions       []Mentions `json:"mentions"`
	UpperMessageID string     `json:"upper_message_id"`
}

type Response struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

type Chats struct {
	Items     []*ChatItem `json:"items"`
	PageToken string      `json:"page_token"`
	HasMore   bool        `json:"has_more"`
}

type ChatItem struct {
	ChatId      string `json:"chat_id"`
	Avatar      string `json:"avatar"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     string `json:"owner_id"`
	OwnerIdType string `json:"owner_id_type"`
}
