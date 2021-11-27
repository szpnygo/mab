package messages

import (
	"strconv"
	"time"
)

type WxMpMessageBody struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        uint64 `xml:"MsgId"`
	PicUrl       string `xml:"PicUrl"`
	MediaID      string `xml:"MediaId"`
	Format       string `xml:"Format"`
	Recognition  string `xml:"Recognition"`
	ThumbMediaID string `xml:"ThumbMediaId"`
	LocationX    string `xml:"Location_X"`
	LocationY    string `xml:"Location_Y"`
	Scale        string `xml:"Scale"`
	Label        string `xml:"Label"`
	Title        string `xml:"Title"`
	Description  string `xml:"Description"`
	Url          string `xml:"Url"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
}

func (msg *WxMpMessageBody) ReplayTextMessage(content string) *TextMessage {
	return &TextMessage{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   strconv.FormatInt(time.Now().UTC().Unix(), 10),
		MsgType:      "text",
		Content:      content,
	}
}
