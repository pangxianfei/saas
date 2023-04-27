package requests

// UserRegister 注册验证器
type UserRegister struct {
	Mobile     string `json:"mobile" validate:"required,len=11"`
	TenantName string `json:"tenant_name" validate:"required"`
	Password   string `json:"password" validate:"required,min=7,max=24"`
	RePassword string `json:"rePassword" validate:"eqfield=Password"`
}
