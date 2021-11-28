package messages

type WxOpenMessageBody struct {
	AppId                        string `xml:"AppId"`
	CreateTime                   string `xml:"CreateTime"`                   //时间戳
	InfoType                     string `xml:"InfoType"`                     //通知类型
	ComponentVerifyTicket        string `xml:"ComponentVerifyTicket"`        //Ticket 内容
	AuthorizerAppid              string `xml:"AuthorizerAppid"`              //公众号或小程序的 appid
	AuthorizationCode            string `xml:"AuthorizationCode"`            //授权码
	AuthorizationCodeExpiredTime string `xml:"AuthorizationCodeExpiredTime"` //授权码过期时间
	PreAuthCode                  string `xml:"PreAuthCode"`                  //预授权码
}
