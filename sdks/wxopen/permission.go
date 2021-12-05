package wxopen

type PermissionType int

// 小程序权限集
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/product/miniprogram_authority.html
const (
	//获取小程序码
	//基于该权限可获取小程序码和小程序二维码
	MINIAPP_GET_CODE PermissionType = 17
	//小程序开发与数据分析
	//基于该权限可进行小程序开发以及代码管理和数据分析析。
	//注意，小程序的开发权限集授权给服务商后，为了避免代码版本互相覆盖 ，小程序则无法再通过mp进行发版本。
	MINIAPP_DEVELOPER PermissionType = 18
	//小程序客服管理
	//基于该权限可实现在小程序中接收和发送客服消息以进行小程序客服消息管理
	MINIAPP_KF PermissionType = 19
	//开放平台帐号管理
	//基于该权限可实现将小程序绑定/解除绑定开放平台帐号
	MINIAPP_OPEN_ACCOUNT PermissionType = 25
	//小程序基本信息管理
	//基于该权限可设置小程序名称、头像、简介、类目等基本信息
	MINIAPP_BASE_INFO PermissionType = 30
	//小程序认证名称检测
	//基于该权限可进行检测小程序认证的名称是否符合规则
	MINIAPP_NAME_CHECK PermissionType = 31
	//微信卡路里管理
	//基于该权限可为小程序提供用户卡路里同步、授权查询、兑换功能
	MINIAPP_CALORIE PermissionType = 36
	//附近的小程序管理
	//基于该权限可为小程序创建附近地点，并设置小程序展示在“附近的小程序”入口中
	MINIAPP_AROUND PermissionType = 37
	//小程序插件管理
	//于该权限可代小程序申请、添加和使用插件并进行管理
	MINIAPP_PLUGIN PermissionType = 40
	//好物圈管理
	//基于该权限将小程序的物品、订单、收藏等信息同步至好物圈
	MINIAPP_GOOD_PRODUCT PermissionType = 41
	//微信物流服务
	//基于该权限可以使用发货、查询组件、消息组件等微信物流服务
	MINIAPP_LOGISTICS PermissionType = 45
	//小程序搜索
	//基于该权限可将小程序页面推送给搜索引擎，增加小程序页面在搜索的收录与曝光机会
	MINIAPP_SEARCH PermissionType = 57
	//小程序广告管理
	//基于该权限可为小程序广告主进行微信广告的投放和管理
	MINIAPP_AD PermissionType = 65
	//获取小程序链接
	//基于该权限可获取小程序URL Scheme、URL Link以及Short Link，从而实现从短信、邮件、微信外网页等场景打开小程序以及在微信内拉起小程序
	MINIAPP_LINK PermissionType = 88
)

//公众号权限集
//https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/product/offical_account_authority.html
const (
	//消息管理权限
	//帮助公众号接收用户消息，进行人工客服回复或自动回复
	OFFICIAL_ACCOUNTS_MESSAGES PermissionType = 1
	//用户管理权限
	//帮助公众号获取用户信息，进行用户管理
	OFFICIAL_ACCOUNTS_USER PermissionType = 2
	//帐号服务权限
	//帮助认证、设置公众号，进行帐号管理
	OFFICIAL_ACCOUNTS_ACCOUNT PermissionType = 3
	//网页服务权限
	//帮助公众号实现第三方网页服务和活动
	OFFICIAL_ACCOUNTS_WEB PermissionType = 4
	//群发与通知权限
	//帮助公众号进行群发和模板消息业务通知
	OFFICIAL_ACCOUNTS_TOAST PermissionType = 7
	//微信扫一扫权限
	//助公众号使用微信扫一扫
	OFFICIAL_ACCOUNTS_SCAN PermissionType = 9
	//素材管理权限
	//帮助公众号管理多媒体素材，用于客服等业务
	OFFICIAL_ACCOUNTS_MATERIAL PermissionType = 11
	//小程序管理权限
	//可新增关联小程序，并对公众号已关联的小程序进行管理
	OFFICIAL_ACCOUNTS_MINIAPP PermissionType = 33
)
