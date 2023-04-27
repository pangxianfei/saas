package requests

// CreateRole 创建角色验证器
type CreateRole struct {
	Name string `json:"name" validate:"required"`
}

// SelectRole 查询角色列表验证器
type SelectRole struct {
	RoleId int64 `json:"role_id" validate:"required,gt=0"`
	UserId int64 `json:"user_id" validate:"required,gt=0"`
}

// QueryRole 查询角色列表验证器
type QueryRole struct {
	RoleId int64 `json:"role_id" validate:"required,gt=0"`
}
