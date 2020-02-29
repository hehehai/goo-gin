package e

//http 响应状态值
const (
	SUCCESS = 200 // 成功
	ERROR = 500 // 响应错误
	INVALID_PARAMS = 400 // 发送值错误

	ERROR_EXIST_TAG = 10001 // 已有标签
	ERROR_NOT_EXIST_TAG = 10002 // 不存在的标签
	ERROR_NOT_EXIST_ARTICLE = 10003 // 不存在的文章

	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001 // token 认证失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002 // token 认证超时
	ERROR_AUTH_TOKEN = 20003 // token 认证错误
	ERROR_AUTH = 20004 // 认证错误
)