package requests

// TenantInstance 角色分配权限验证器
type TenantInstance struct {
	UserId int64 `json:"UserId" validate:"required,gt=0"`
}
