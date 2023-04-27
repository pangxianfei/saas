package permission

import (
	"gitee.com/pangxianfei/saas/requests"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
	"github.com/kataras/iris/v12"
)

type Permission interface {
	// AddPermissionToApp 平台级
	AddPermissionToApp(cxt iris.Context, appId int64) (err error) //*添加整个应用所有权限

	GetAppMenu(cxt iris.Context, appId int64) []sysmdel.Permissions
	SyncPermissionsTo(cxt iris.Context, permission requests.SyncPermission) (err error) //*撤销权限、并添加新的权限
	// RevokePermissionTo 应用级
	RevokePermissionTo(cxt iris.Context, permission requests.SyncPermission) (err error) //*撤销用户权限

	AddRole(cxt iris.Context, role requests.CreateRole) (*sysmdel.SysRoles, error)              //*添加角色
	RemoveRole(cxt iris.Context, roleId int64) (err error)                                      //*删除角色
	HasAnyRole(cxt iris.Context, SelectCreate requests.QueryRole) interface{}                   //*用户角色列表
	SyncPermissionToRoles(cxt iris.Context, permission requests.RolePermission) (err error)     //*给角色同步权限
	GiveRoleToPermission(cxt iris.Context, giveRolePermission requests.GiveRolePermission) bool //*给角色添加一个权限
	GiveUserRolePermission(cxt iris.Context, selectRole requests.SelectRole) error              //*给用户添加角色权限
	HasPermission(cxt iris.Context, routeName string) bool

	RemovePermissionToRoles(cxt iris.Context, permission requests.GiveRolePermission) (err error) //*删除角色中某个权限
	HasRole(cxt iris.Context, roleId int64) bool                                                  //*是否具有某个角色

	RemoveUserRole(cxt iris.Context, selectRole requests.SelectRole) bool                      //取消用户角色权限
	HasRoleToPermission(cxt iris.Context, giveRolePermission requests.GiveRolePermission) bool //*确定角色是否具有某种权限
	RevokeRoleToPermission(cxt iris.Context, selectRole requests.SelectRole) bool              //*撤销角色权限

}
