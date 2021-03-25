package shanyan

const (
	MSG_SUCCESS 			= 200000    // 请求成功
	MSG_PARAMS_ERROR 		= 400001    // 参数校验异常
	MSG_SIGN_ERROR 			= 403000    // 用户校验失败
	MSG_transfer_ERROR 		= 415000    // 请求数据转换异常
	MSG_SYSTEM_ERROR 		= 500000    // 系统异常
	MSG_HANDLE_ERROR 		= 500002    // 数据处理异常
	MSG_ACTION_ERROR 		= 500003    // 业务操作失败
	MSG_CALL_ERROR 			= 500004    // 远程调用失败
	MSG_BALANCE_ERROR 	    = 500005    // 账户余额异常
	MSG_EXTERNAL_ERROR 		= 500006    // 请求外部系统失败
	MSG_SYSTEM_TIMEOUT		= 504000    // 系统超时
	MSG_MERCHANT_WITHOUT 	= 400101    // 在下游系统中的商户信息不存在
	MSG_ACCOUNT_DISABLE 	= 403101    // 账户被下游系统禁用
	MSG_ACCOUNT_NOTACTIVE   = 403102    // 账户在下游系统中没有被激活
	MSG_ACCOUNT_INSUFFICIENT_QUANTITY = 510101 // 在下游系统中的用户产品可用数量不足
	MSG_IP_DISABLE  		= 510101 	// 商户IP地址在下游系统中不合法
	MSG_BLACKLIST 			= 400200 	// 黑名单列表
	MSG_PHONE_EMPTY         = 400201    // 手机号码不能为空
	MSG_ACCOUNT_EMPTY       = 400901    // 账户信息不存在
	MSG_TYPE_EMPTY          = 400902    // 应用类型信息不存在
	MSG_EMAIL_NOT_SETTING   = 500901    // 邮箱未设置
	MSG_ACCOUNT_EXISTS      = 500902    // 账户信息已存在
	MSG_ACCOUNT_RELEVANT    = 500903    // 账户相关能力已激活
)

type MobileQueryReqBody struct {
	AppID  			string
	Token  			string
	Sign   			string
}

type BaseResponse struct {
	Code    string 	`json:"code"`
	Message string  `json:"message"`
	chargeStatus int `json:"chargeStatus"`
}

type MobileQueryResponse struct {
	BaseResponse *BaseResponse
	Data struct{
		MobileName string 	`json:"mobileName"`
		TradeNo    string 	`json:"tradeNo"`
	} `json:"data"`
}