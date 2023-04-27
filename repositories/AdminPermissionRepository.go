package repositories

import (
	"errors"

	"gitee.com/pangxianfei/frame/simple/sqlcmd"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
	"gorm.io/gorm"
)

var AdminPermissionRepository = new(AdminPermissionDao)

type AdminPermissionDao struct {
}

// Delete 主键条件删除
func (r *AdminPermissionDao) Delete(db *gorm.DB, id int64) error {
	return db.Delete(&sysmdel.AdminPermissions{}, "id = ?", id).Error
}

func (r *AdminPermissionDao) Find(db *gorm.DB, cnd *sqlcmd.Cnd) (list []sysmdel.AdminPermissions) {
	cnd.Find(db, &list)
	return
}

func (r *AdminPermissionDao) Create(db *gorm.DB, TenantId int64, UserId int64, permission []int64) (createErr error) {
	r.IsHasTable(db)
	if UserId > 0 && len(permission) > 0 {
		for _, permissionItem := range permission {
			rowsAffected := db.Find(&sysmdel.AdminPermissions{}, sysmdel.AdminPermissions{UserId: UserId, PermissionId: permissionItem}).RowsAffected
			if rowsAffected <= 0 {
				SysPermissions := PermissionRepository.Take(db, &sysmdel.Permissions{PermissionId: permissionItem})
				AdminPermissions := &sysmdel.AdminPermissions{UserId: UserId, PermissionId: permissionItem, AppId: SysPermissions.AppId, RouteName: SysPermissions.RouteName, RouteNameMd5Value: SysPermissions.RouteNameMd5Value, TenantsId: TenantId}
				// 通过数据的指针来创建
				if createErr = db.Create(AdminPermissions).Error; createErr != nil {
					return createErr
				}
			}
		}
	}
	return createErr
}

// Revoke 删除用户权限
func (r *AdminPermissionDao) Revoke(db *gorm.DB, UserId int64, permission []int64) error {
	if UserId > 0 && len(permission) > 0 {
		cmd := sqlcmd.NewCnd().Where("user_id", UserId).In("permission_id", permission)
		AdminPermission := r.Find(db, cmd)
		var deleteErr error
		if len(AdminPermission) > 0 {
			for _, AdminPermissionItem := range AdminPermission {
				deleteErr = r.Delete(db, AdminPermissionItem.Id)
				if deleteErr = r.Delete(db, AdminPermissionItem.Id); deleteErr != nil {
					return deleteErr
				}
			}
			return deleteErr
		}
	}
	return errors.New("条件不足,删除失败")
}

func (r *AdminPermissionDao) IsHasTable(db *gorm.DB) {
	if db.Migrator().HasTable(&sysmdel.AdminPermissions{}) == false {
		db.Migrator().CreateTable(&sysmdel.AdminPermissions{})
	}
}

func (r *AdminPermissionDao) RevokeRoleToPermission(db *gorm.DB, roleId int64) error {
	if err := db.Where("role_id = ?", roleId).Delete(&sysmdel.AdminPermissions{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *AdminPermissionDao) DeleteUserPermission(db *gorm.DB, userId int64, roleId int64) error {
	if err := db.Where(&sysmdel.AdminPermissions{UserId: userId, RoleId: roleId}).Delete(&sysmdel.AdminPermissions{}).Error; err != nil {
		return err
	}
	return nil
}
