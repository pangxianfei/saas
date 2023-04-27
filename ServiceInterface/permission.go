package ServiceInterface

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"

	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/saas/sysmodel"
)

type Permission interface {
	GetTenantId(cxt iris.Context) int64
	GetUserDb(tenantId int64) *gorm.DB
	AddPermissionToApp(cxt iris.Context, appId int64) (err error)
	// AddRole 添加角色
	AddRole(cxt iris.Context, role requests.CreateRole) (*sysmdel.SysRoles, error)
	// SyncPermissionToRoles 添加角色权限
	SyncPermissionToRoles(cxt iris.Context, permission requests.RolePermission) error
	// HasAnyRole 角色权限列表
	HasAnyRole(cxt iris.Context, selectCreate requests.QueryRole) interface{}
	RemovePermissionToRoles(cxt iris.Context, permission requests.GiveRolePermission) error
	SyncPermissionsTo(cxt iris.Context, permission requests.SyncPermission) error
	RevokePermissionTo(cxt iris.Context, UserId int64, permission []int64) error
	// RemoveRole 删除角色
	RemoveRole(cxt iris.Context, roleId int64) error
	GiveRoleToPermission(cxt iris.Context, giveRolePermission requests.GiveRolePermission) bool
	GiveUserRolePermission(cxt iris.Context, roleId int64, userId int64) error
	HasRole(cxt iris.Context, roleId int64) bool
	HasPermission(cxt iris.Context, routeName string) bool
	RevokeRoleToPermission(cxt iris.Context, roleId int64) bool
	HasRoleToPermission(cxt iris.Context, giveRolePermission requests.GiveRolePermission) bool
	RemoveUserRole(cxt iris.Context, SelectRole requests.SelectRole) bool
	GetAppMenu(cxt iris.Context, appId int64) ([]sysmdel.Permissions, error)
}
