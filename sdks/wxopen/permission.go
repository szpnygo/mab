package wxopen

type PermissionType int

// 小程序权限集
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
)
