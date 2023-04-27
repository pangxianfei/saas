package permission

import (
	"github.com/kataras/iris/v12"

	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/saas/services"
	"gitee.com/pangxianfei/saas/sysmodel"
)

type GatePermission struct {
	TenantId int64
}

func (p *GatePermission) GetAppMenu(cxt iris.Context, appId int64) []sysmdel.Permissions {
	appMenu, err := services.PermissionService.GetAppMenu(cxt, appId)
	if err != nil {
		return nil
	}
	return appMenu
}

// AddPermissionToApp 平台级
func (p *GatePermission) AddPermissionToApp(cxt iris.Context, appId int64) (err error) {
	return services.PermissionService.AddPermissionToApp(cxt, appId)
}

// SyncPermissionsTo 撤销权限、并添加新的权限
func (p *GatePermission) SyncPermissionsTo(cxt iris.Context, permission requests.SyncPermission) (err error) {
	return services.PermissionService.SyncPermissionsTo(cxt, permission)
}

// RevokePermissionTo 撤销用户权限
func (p *GatePermission) RevokePermissionTo(cxt iris.Context, permission requests.SyncPermission) (err error) {
	return services.PermissionService.RevokePermissionTo(cxt, permission.UserId, permission.Permission)
}

func (p *GatePermission) AddRole(cxt iris.Context, role requests.CreateRole) (*sysmdel.SysRoles, error) {
	return services.PermissionService.AddRole(cxt, role)
}

// RemoveRole 删除角色
func (p *GatePermission) RemoveRole(cxt iris.Context, roleId int64) (err error) {
	return services.PermissionService.RemoveRole(cxt, roleId)
}

// SyncPermissionToRoles 给角色同步权限
func (p *GatePermission) SyncPermissionToRoles(cxt iris.Context, permission requests.RolePermission) error {
	return services.PermissionService.SyncPermissionToRoles(cxt, permission)
}

// GiveRoleToPermission 给角色添加一个权限
func (p *GatePermission) GiveRoleToPermission(cxt iris.Context, giveRolePermission requests.GiveRolePermission) bool {
	return services.PermissionService.GiveRoleToPermission(cxt, giveRolePermission)
}

// RemovePermissionToRoles 删除角色中某个权限
func (p *GatePermission) RemovePermissionToRoles(cxt iris.Context, permission requests.GiveRolePermission) error {
	return services.PermissionService.RemovePermissionToRoles(cxt, permission)
}

// HasRole 是否具有某个角色
func (p *GatePermission) HasRole(cxt iris.Context, roleId int64) bool {
	return services.PermissionService.HasRole(cxt, roleId)
}

func (p *GatePermission) RemoveUserRole(cxt iris.Context, SelectRole requests.SelectRole) bool {
	return services.PermissionService.RemoveUserRole(cxt, SelectRole)
}

// HasAnyRole 用户角色列表
func (p *GatePermission) HasAnyRole(cxt iris.Context, QueryRole requests.QueryRole) interface{} {
	return services.PermissionService.HasAnyRole(cxt, QueryRole)
}

// GiveUserRolePermission 给用户添加角色权限
func (p *GatePermission) GiveUserRolePermission(cxt iris.Context, selectRole requests.SelectRole) error {
	return services.PermissionService.GiveUserRolePermission(cxt, selectRole.RoleId, selectRole.UserId)
}

// HasRoleToPermission 确定角色是否具有某种权限
func (p *GatePermission) HasRoleToPermission(cxt iris.Context, giveRolePermission requests.GiveRolePermission) bool {
	return services.PermissionService.HasRoleToPermission(cxt, giveRolePermission)
}

// RevokeRoleToPermission 撤销角色权限
func (p *GatePermission) RevokeRoleToPermission(cxt iris.Context, selectRole requests.SelectRole) bool {
	return services.PermissionService.RevokeRoleToPermission(cxt, selectRole.RoleId)
}

// HasPermission 给用户添加角色权限
func (p *GatePermission) HasPermission(cxt iris.Context, routeName string) bool {
	return services.PermissionService.HasPermission(cxt, routeName)
}
