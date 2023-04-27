package repositories

import (
	"errors"
	"gitee.com/pangxianfei/saas/sysmodel"
	"gorm.io/gorm"
)

var UserRoleRepository = new(UserRoleDao)

type UserRoleDao struct {
}

func (r *UserRoleDao) IsHasTable(db *gorm.DB) {
	if db.Migrator().HasTable(&sysmdel.UserRoles{}) == false {
		db.Migrator().CreateTable(&sysmdel.UserRoles{})
		db.Migrator().CreateTable(&sysmdel.AdminPermissions{})
	}
}

// Delete 主键条件删除
func (r *UserRoleDao) Delete(db *gorm.DB, id int64) error {
	return db.Delete(&sysmdel.UserRoles{}, "id = ?", id).Error
}

func (r *UserRoleDao) HasRole(db *gorm.DB, userId int64, roleId int64) bool {
	roleInfo := r.Take(db, "user_id = ? and role_id = ?", userId, roleId)
	if roleInfo.UserId > 0 && roleInfo.RoleId > 0 && roleInfo.Id > 0 {
		return true
	}
	return false
}
func (r *UserRoleDao) Take(db *gorm.DB, where ...interface{}) *sysmdel.UserRoles {
	ret := &sysmdel.UserRoles{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *UserRoleDao) Create(db *gorm.DB, tenantId int64, userId int64, roleId int64) (createErr error) {
	r.IsHasTable(db)
	rolesInfo := RolesRepository.RoleInfo(db, roleId)
	if rolesInfo.Id > 0 && rolesInfo.Permissions != nil && len(rolesInfo.Permissions) > 0 {
		UserRoles := &sysmdel.UserRoles{UserId: userId, RoleId: roleId, TenantsId: tenantId}
		// 通过数据的指针来创建
		if createErr = db.FirstOrCreate(UserRoles, UserRoles).Error; createErr == nil {
			for _, permissionItem := range rolesInfo.Permissions {
				SysPermissions := PermissionRepository.Take(db, &sysmdel.Permissions{PermissionId: permissionItem.PermissionId})
				AdminPermissions := &sysmdel.AdminPermissions{
					UserId:            userId,
					PermissionId:      SysPermissions.PermissionId,
					AppId:             SysPermissions.AppId,
					RouteName:         SysPermissions.RouteName,
					RouteNameMd5Value: SysPermissions.RouteNameMd5Value,
					RoleId:            UserRoles.RoleId,
					TenantsId:         UserRoles.TenantsId,
				}
				db.FirstOrCreate(AdminPermissions, AdminPermissions)
			}
			return
		}
	}
	return errors.New("角色未授权")
}

func (r *UserRoleDao) RemoveUserRole(db *gorm.DB, userId int64, roleId int64) error {
	if err := db.Where(&sysmdel.UserRoles{UserId: userId, RoleId: roleId}).Delete(&sysmdel.UserRoles{}).Error; err != nil {
		return err
	}
	return nil
}
