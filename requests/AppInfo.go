package requests

// AppInfo 应用验证器
type AppInfo struct {
	Key         string `json:"key" validate:"required"`
	AppName     string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
