package requests

// RolePermission 角色分配权限验证器
type RolePermission struct {
	RoleId     int64   `json:"role_id" validate:"required,gt=0"`
	Permission []int64 `json:"permission" validate:"required"`
}

// GiveRolePermission 角色分配权限验证器
type GiveRolePermission struct {
	RoleId       int64 `json:"role_id" validate:"required,gt=0" field_error_info:"用户名最少6个字符"`
	PermissionId int64 `json:"permission_id" validate:"required,gt=0" field_error_info:"用户名最少6个字符"`
}

// SyncPermission 撤销权限、并添加新的权限验证器
type SyncPermission struct {
	Permission []int64 `json:"permission" validate:"required"`
	UserId     int64   `json:"user_id" validate:"required"`
}
