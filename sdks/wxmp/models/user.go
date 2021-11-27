package models

type UserInfo struct {
	OpenID         string `json:"openid"`
	Subscribe      int    `json:"subscribe"`
	NickName       string `json:"nickname"`
	Language       string `json:"language"`
	HeadImgURL     string `json:"headimgurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	UnionID        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupID        int    `json:"groupid"`
	TagIDList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QRScene        int    `json:"qr_scene"`
	QRSceneStr     string `json:"qr_scene_str"`
}
