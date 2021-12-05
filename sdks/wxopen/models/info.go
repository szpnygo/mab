package models

type GetAuthInfoRequest struct {
	ComponentAppID  string `json:"component_appid"`
	AuthorizerAppID string `json:"authorizer_appid"`
}

type AuthInfoResult struct {
	BaseResponse
	AuthorizerInfo    AuthorizerInfo    `json:"authorizer_info"`
	AuthorizationInfo AuthorizationInfo `json:"authorization_info"`
}

type AuthorizerInfo struct {
	NickName        string           `json:"nick_name"`         // 授权方昵称
	HeadImg         string           `json:"head_img"`          // 授权方头像
	ServiceTypeInfo ServiceTypeInfo  `json:"service_type_info"` // 授权方小程序/公众号类型
	VerifyTypeInfo  VerifyTypeInfo   `json:"verify_type_info"`  // 授权方认证类型
	UserName        string           `json:"user_name"`         // 授权方公众号的原始ID
	PrincipalName   string           `json:"principal_name"`    //主体名称
	BusinessInfo    BusinessInfo     `json:"business_info"`     //用以了解功能的开通状况
	Alias           string           `json:"alias"`             //公众号所设置的微信号，可能为空
	QrcodeURL       string           `json:"qrcode_url"`        //二维码图片的 URL，开发者最好自行也进行保存
	Signature       string           `json:"signature"`         //帐号介绍
	MiniProgramInfo *MiniProgramInfo `json:"MiniProgramInfo"`   //小程序配置，根据这个字段判断是否为小程序类型授权
}

func (a *AuthorizerInfo) IsMiniProgram() bool {
	return a.MiniProgramInfo != nil
}

type ServiceTypeInfo struct {
	//0代表订阅号、普通小程序
	//1代表由历史老帐号升级后的订阅号
	//2代表服务号
	//12试用小程序
	//4小游戏
	//10小商店
	//2或者3门店小程序
	ID int `json:"id"`
}
type VerifyTypeInfo struct {
	//-1未认证
	//0微信认证
	//1新浪微博认证
	//2腾讯微博认证
	//3已资质认证通过但还未通过名称认证
	//4已资质认证通过、还未通过名称认证，但通过了新浪微博认证
	//5已资质认证通过、还未通过名称认证，但通过了腾讯微博认证
	ID int `json:"id"`
}
type BusinessInfo struct {
	OpenStore int `json:"open_store"` //是否开通微信门店功能
	OpenScan  int `json:"open_scan"`  //是否开通微信扫商品功能
	OpenPay   int `json:"open_pay"`   //是否开通微信支付功能
	OpenCard  int `json:"open_card"`  //是否开通微信卡券功能
	OpenShake int `json:"open_shake"` //是否开通微信摇一摇功能
}

type MiniProgramInfo struct {
	Network     Network      `json:"network"`    //小程序配置的合法域名信息
	Categories  []Categories `json:"categories"` //小程序配置的类目信息
	VisitStatus int          `json:"visit_status"`
}

type Network struct {
	RequestDomain   []string      `json:"RequestDomain"`
	WsRequestDomain []interface{} `json:"WsRequestDomain"`
	UploadDomain    []interface{} `json:"UploadDomain"`
	DownloadDomain  []interface{} `json:"DownloadDomain"`
	BizDomain       []interface{} `json:"BizDomain"`
	UDPDomain       []interface{} `json:"UDPDomain"`
}
type Categories struct {
	First  string `json:"first"`
	Second string `json:"second"`
}
