package response

// ErrorModel 错误返回模型
type ErrorModel struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"Data"`
}

// ErrorUnregisteredTenantAppDb 403-未没注册的应用
func ErrorUnregisteredTenantAppDb() ErrorModel {
	return ErrorModel{
		Code:    403,
		Message: "未注册的应用或者数据库连接失败",
		Success: false,
		Data:    "未注册的应用或者数据库连接失败",
	}
}

func Error(data interface{}) ErrorModel {
	return ErrorModel{
		Code:    403,
		Message: "没有权限",
		Success: false,
		Data:    data,
	}
}
func ErrorNoHaveAuthority() ErrorModel {
	return ErrorModel{
		Code:    403,
		Message: "没有权限",
		Success: false,
		Data:    "没有权限",
	}
}

// ErrorTokenInvalidation 401-未认证登录
func ErrorTokenInvalidation() ErrorModel {
	return ErrorModel{
		Code:    401,
		Message: "令牌失效",
		Success: false,
		Data:    "令牌失效",
	}
}

// ErrorUnauthorized 401-未认证登录
func ErrorUnauthorized() ErrorModel {
	return ErrorModel{
		Code:    401,
		Message: "未认证登录",
		Success: false,
		Data:    "未认证登录",
	}
}
