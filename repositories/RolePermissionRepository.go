package repositories

import (
	"errors"

	"gitee.com/pangxianfei/frame/kernel/debug"
	"gitee.com/pangxianfei/frame/simple/sqlcmd"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
	"gorm.io/gorm"
)

var RolePermissionRepository = new(RoleHasPermissionDao)

type RoleHasPermissionDao struct {
}

// Delete 主键条件删除
func (r *RoleHasPermissionDao) Delete(db *gorm.DB, id int64) error {
	return db.Delete(&sysmdel.RolePermissions{}, "id = ?", id).Error
}

// DeleteRolePermission 角色ID条件删除
func (r *RoleHasPermissionDao) DeleteRolePermission(db *gorm.DB, roleId int64) error {
	return db.Where("role_id = ?", roleId).Delete(&sysmdel.RolePermissions{}).Error
}

func (r *RoleHasPermissionDao) Take(db *gorm.DB, where ...interface{}) *sysmdel.RolePermissions {
	ret := &sysmdel.RolePermissions{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

// Find 返回权限列表(菜单)
func (r *RoleHasPermissionDao) Find(db *gorm.DB, cnd *sqlcmd.Cnd) []sysmdel.Permissions {
	return PermissionRepository.Find(db, cnd)
}

// GetByRoleId 查询角色下是否分配有权限
func (r *RoleHasPermissionDao) GetByRoleId(db *gorm.DB, roleId int64) *sysmdel.RolePermissions {
	return r.Take(db, "role_id = ?", roleId)
}

func (r *RoleHasPermissionDao) GetById(db *gorm.DB, id int64) *sysmdel.RolePermissions {
	return r.Take(db, "id = ?", id)
}

// CreateOne 添加一个角色权限
func (r *RoleHasPermissionDao) CreateOne(db *gorm.DB, tenantId int64, roleId int64, permissionId int64) (err error) {
	//表不存则创建表
	r.IsHasTable(db)

	SysPermissions := PermissionRepository.Take(db, "permission_id =?", permissionId)

	if SysPermissions == nil {
		return nil
	}

	return db.Transaction(func(tx *gorm.DB) error {
		newPermission := &sysmdel.RolePermissions{RoleId: roleId, PermissionId: SysPermissions.PermissionId, TenantsId: tenantId, AppId: SysPermissions.AppId}
		if err := tx.Create(newPermission).Error; err != nil {
			return err
		}
		return nil
	})

}

// Create 创建
func (r *RoleHasPermissionDao) Create(db *gorm.DB, tenantId int64, roleId int64, permission []int64) error {
	//表不存则创建表
	r.IsHasTable(db)
	db.Where(&sysmdel.RolePermissions{RoleId: roleId}).Delete(&sysmdel.RolePermissions{})
	cmd := sqlcmd.NewCnd().In("permission_id", permission)
	SysPermissions := r.Find(db, cmd)

	if len(SysPermissions) <= 0 {
		return nil
	}
	var err error
	for _, permissionItem := range SysPermissions {

		err = db.Transaction(func(tx *gorm.DB) error {
			newPermission := &sysmdel.RolePermissions{RoleId: roleId, PermissionId: permissionItem.PermissionId, TenantsId: tenantId, AppId: permissionItem.AppId}
			if err := tx.Create(newPermission).Error; err != nil {
				return err
			}
			return nil
		})

	}

	return err
}

func (r *RoleHasPermissionDao) RemovePermission(db *gorm.DB, roleId int64, PermissionId int64) error {
	if roleId > 0 || PermissionId > 0 {
		return db.Delete(&sysmdel.RolePermissions{RoleId: roleId, PermissionId: PermissionId}, PermissionId).Error
	}
	return errors.New("参数必需")
}

func (r *RoleHasPermissionDao) IsHasTable(db *gorm.DB) {
	if db.Migrator().HasTable(&sysmdel.RolePermissions{}) == false {
		db.Migrator().CreateTable(&sysmdel.RolePermissions{})
	}
}

func (r *RoleHasPermissionDao) RoleToPermission(db *gorm.DB, roleId int64, PermissionId int64) bool {
	whereModel := &sysmdel.RolePermissions{
		RoleId:       roleId,
		PermissionId: PermissionId,
	}
	Result := db.Where(whereModel).First(&sysmdel.RolePermissions{})
	if Result.Error == nil && Result.RowsAffected > 0 {
		debug.Dd("yes")
		return true
	}
	debug.Dd("NO")
	return false
}
