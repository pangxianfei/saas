package repositories

import (
	"errors"

	"gitee.com/pangxianfei/frame/simple"
	"gitee.com/pangxianfei/frame/simple/sqlcmd"
	"gorm.io/gorm"

	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

var PermissionRepository = new(synPermissionDao)

type synPermissionDao struct {
}

// DeleteRolePermission 角色ID条件删除
func (r *synPermissionDao) DeleteRolePermission(db *gorm.DB, roleId int64) error {
	return db.Where("role_id = ?", roleId).Delete(&sysmdel.RolePermissions{}).Error
}
func (r *synPermissionDao) Take(db *gorm.DB, where ...interface{}) *sysmdel.Permissions {
	ret := &sysmdel.Permissions{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

// Find 返回列表
func (r *synPermissionDao) Find(db *gorm.DB, cnd *sqlcmd.Cnd) (list []sysmdel.Permissions) {
	cnd.Find(db, &list)
	return
}

func (r *synPermissionDao) GetByTenantName(TenantName string) *sysmdel.Permissions {
	return r.Take(simple.DB(), "tenant_name = ?", TenantName)
}

func (r *synPermissionDao) Create(db *gorm.DB, SysPermissions *sysmdel.Permissions) (Tenants *sysmdel.Permissions, err error) {
	r.IsHasTable(db)
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.FirstOrCreate(SysPermissions, sysmdel.Permissions{PermissionId: SysPermissions.PermissionId}).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return SysPermissions, err
}

func (r *synPermissionDao) HasPermission(db *gorm.DB, userId int64, routeName string) *sysmdel.AdminPermissions {
	Permissions := &sysmdel.AdminPermissions{UserId: userId, RouteName: routeName}
	PermissionsRows := db.Where(Permissions).First(&Permissions)
	if PermissionsRows.Error != nil {
		return nil
	}
	if PermissionsRows.RowsAffected > 0 {
		return Permissions
	}
	return nil

}

func (r *synPermissionDao) GetAppMenu(db *gorm.DB, appId int64) ([]sysmdel.Permissions, error) {
	var Permission []sysmdel.Permissions
	if err := db.Debug().Preload("Children", "is_menu = ?", 1).Where(&sysmdel.Permissions{AppId: appId, IsMenu: 1}).Where("pid = ?", 0).Find(&Permission).Error; err != nil {
		return Permission, errors.New("无记录")
	}
	return Permission, nil
}

func (r *synPermissionDao) IsHasTable(db *gorm.DB) {
	if db.Migrator().HasTable(&sysmdel.Permissions{}) == false {
		_ = db.Migrator().CreateTable(&sysmdel.Permissions{})
	}
}
