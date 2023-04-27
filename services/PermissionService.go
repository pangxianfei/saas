package services

import (
	"errors"

	"gitee.com/pangxianfei/frame/simple"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"

	"gitee.com/pangxianfei/saas/buffer"
	"gitee.com/pangxianfei/saas/repositories"
	"gitee.com/pangxianfei/saas/requests"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

var PermissionService = new(synPermissionService)

type synPermissionService struct {
	TenantId int64
}

func (c *synPermissionService) GetTenantId(cxt iris.Context) int64 {
	c.TenantId = InstanceService.GetTenantId(cxt)
	return c.TenantId
}

func (c *synPermissionService) GetUserDb(tenantId int64) *gorm.DB {
	if userDb, err := InstanceService.GetTenantUserDb(tenantId); err == nil {
		return userDb
	}
	return nil
}

func (c *synPermissionService) AddPermissionToApp(cxt iris.Context, appId int64) (err error) {
	c.GetTenantId(cxt)
	AuthorityList := repositories.AuthorityRepository.GetByAppIdList(simple.DB(), appId)

	if len(AuthorityList) > 0 {
		userDb := c.GetUserDb(c.TenantId)
		for _, AuthorityItem := range AuthorityList {
			createSysPermissions := &sysmdel.Permissions{
				PermissionId: AuthorityItem.Id,
				Pid:          AuthorityItem.Pid,
				AppId:        AuthorityItem.AppId,
				Name:         AuthorityItem.Name,
				MenuName:     AuthorityItem.MenuName,
				Description:  AuthorityItem.Description,
				//RegisterFileName:  AuthorityItem.RegisterFileName,
				//MainHandlerName:   AuthorityItem.MainHandlerName,
				Method:        AuthorityItem.Method,
				FormattedPath: AuthorityItem.FormattedPath,
				StaticPath:    AuthorityItem.StaticPath,
				Path:          AuthorityItem.Path,
				//SourceFileName:    AuthorityItem.SourceFileName,
				RouteName:         AuthorityItem.RouteName,
				Status:            AuthorityItem.Status,
				IsMenu:            AuthorityItem.IsMenu,
				Md5Value:          AuthorityItem.Md5Value,
				RouteNameMd5Value: AuthorityItem.RouteNameMd5Value,
			}
			go func() {
				_, err = repositories.PermissionRepository.Create(userDb, createSysPermissions)
			}()
		}
	}
	return err
}

// AddRole 添加角色
func (c *synPermissionService) AddRole(cxt iris.Context, role requests.CreateRole) (*sysmdel.SysRoles, error) {
	c.GetTenantId(cxt)
	if sysRoles, err := repositories.RolesRepository.Create(c.GetUserDb(c.TenantId), &sysmdel.SysRoles{Name: role.Name, TenantsId: c.TenantId}); err == nil {
		return sysRoles, nil
	}
	return nil, nil
}

// SyncPermissionToRoles 添加角色权限
func (c *synPermissionService) SyncPermissionToRoles(cxt iris.Context, permission requests.RolePermission) error {
	c.GetTenantId(cxt)
	if repositories.RolesRepository.GetById(c.GetUserDb(c.TenantId), permission.RoleId) == nil {
		return errors.New("角色不存在")
	}
	return repositories.RolePermissionRepository.Create(c.GetUserDb(c.TenantId), c.TenantId, permission.RoleId, permission.Permission)
}

// HasAnyRole 角色权限列表
func (c *synPermissionService) HasAnyRole(cxt iris.Context, selectCreate requests.QueryRole) interface{} {
	c.GetTenantId(cxt)
	return repositories.RolesRepository.AnyRole(c.GetUserDb(c.TenantId), selectCreate.RoleId)
}

func (c *synPermissionService) RemovePermissionToRoles(cxt iris.Context, permission requests.GiveRolePermission) error {
	c.GetTenantId(cxt)
	if err := repositories.RolePermissionRepository.RemovePermission(c.GetUserDb(c.TenantId), permission.RoleId, permission.PermissionId); err != nil {
		return errors.New("删除失败")
	}
	return nil
}

func (c *synPermissionService) SyncPermissionsTo(cxt iris.Context, permission requests.SyncPermission) error {
	c.GetTenantId(cxt)
	return repositories.AdminPermissionRepository.Create(c.GetUserDb(c.TenantId), c.TenantId, permission.UserId, permission.Permission)
}

func (c *synPermissionService) RevokePermissionTo(cxt iris.Context, UserId int64, permission []int64) error {
	c.GetTenantId(cxt)
	return repositories.AdminPermissionRepository.Revoke(c.GetUserDb(c.TenantId), UserId, permission)
}

// RemoveRole 删除角色
func (c *synPermissionService) RemoveRole(cxt iris.Context, roleId int64) error {
	c.GetTenantId(cxt)
	if hasRoleId := repositories.RolePermissionRepository.GetByRoleId(c.GetUserDb(c.TenantId), roleId); hasRoleId != nil {
		return errors.New("该角色使用中")
	}
	err := repositories.RolesRepository.Delete(c.GetUserDb(c.TenantId), roleId)
	if err == nil {
		return repositories.RolePermissionRepository.DeleteRolePermission(c.GetUserDb(c.TenantId), roleId)
	}

	return err
}

func (c *synPermissionService) GiveRoleToPermission(cxt iris.Context, giveRolePermission requests.GiveRolePermission) bool {
	c.GetTenantId(cxt)
	var permission []int64
	permission = append(permission, giveRolePermission.PermissionId)
	if err := repositories.RolePermissionRepository.CreateOne(c.GetUserDb(c.TenantId), c.TenantId, giveRolePermission.RoleId, giveRolePermission.PermissionId); err != nil {
		return false
	}
	return true
}

func (c *synPermissionService) GiveUserRolePermission(cxt iris.Context, roleId int64, userId int64) error {
	c.GetTenantId(cxt)
	if err := repositories.UserRoleRepository.Create(c.GetUserDb(c.TenantId), c.TenantId, userId, roleId); err != nil {
		return err
	}
	return nil
}

func (c *synPermissionService) HasRole(cxt iris.Context, roleId int64) bool {
	c.GetTenantId(cxt)
	UserId := UserService.GetUserId(cxt)
	return repositories.UserRoleRepository.HasRole(c.GetUserDb(c.TenantId), UserId, roleId)
}

func (c *synPermissionService) HasPermission(cxt iris.Context, routeName string) bool {
	c.GetTenantId(cxt)
	UserId := UserService.GetUserId(cxt)
	if buffer.CachePermission.GetCachePermission(UserId, routeName) {
		return true
	}

	adminPermissions := repositories.PermissionRepository.HasPermission(c.GetUserDb(c.TenantId), UserId, routeName)

	if adminPermissions != nil && adminPermissions.RouteName == routeName && adminPermissions.UserId == UserId {
		//缓存权限
		buffer.CachePermission.SetCachePermission(adminPermissions)
		return true
	}

	return false
}

func (c *synPermissionService) RevokeRoleToPermission(cxt iris.Context, roleId int64) bool {
	c.GetTenantId(cxt)
	UserDb := c.GetUserDb(c.TenantId)

	deleteErr := UserDb.Transaction(func(tx *gorm.DB) error {
		if err := repositories.AdminPermissionRepository.RevokeRoleToPermission(tx, roleId); err != nil {
			return err
		}
		if err := repositories.RolePermissionRepository.DeleteRolePermission(tx, roleId); err != nil {
			return err
		}
		return nil
	})

	return deleteErr == nil
}

func (c *synPermissionService) HasRoleToPermission(cxt iris.Context, giveRolePermission requests.GiveRolePermission) bool {
	c.GetTenantId(cxt)
	return repositories.RolePermissionRepository.RoleToPermission(c.GetUserDb(c.TenantId), giveRolePermission.RoleId, giveRolePermission.PermissionId)
}

func (c *synPermissionService) RemoveUserRole(cxt iris.Context, SelectRole requests.SelectRole) bool {
	c.GetTenantId(cxt)
	UserDb := c.GetUserDb(c.TenantId)
	UserId := UserService.GetUserId(cxt)

	deleteErr := UserDb.Transaction(func(tx *gorm.DB) error {
		if err := repositories.UserRoleRepository.RemoveUserRole(tx, UserId, SelectRole.RoleId); err != nil {
			return err
		}

		if err := repositories.AdminPermissionRepository.DeleteUserPermission(tx, UserId, SelectRole.RoleId); err != nil {
			return err
		}
		return nil
	})

	return deleteErr == nil
}

func (c *synPermissionService) GetAppMenu(cxt iris.Context, appId int64) ([]sysmdel.Permissions, error) {
	c.GetTenantId(cxt)
	UserDb := c.GetUserDb(c.TenantId)
	if Permissions, err := repositories.PermissionRepository.GetAppMenu(UserDb, appId); err == nil {
		return Permissions, err
	}
	return nil, nil
}
