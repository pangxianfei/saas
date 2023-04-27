package repositories

import (
	"errors"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
	"gorm.io/gorm"
)

var RolesRepository = new(RolesDao)

type RolesDao struct {
}

// Delete 主键条件删除
func (r *RolesDao) Delete(db *gorm.DB, id int64) error {
	return db.Delete(&sysmdel.SysRoles{}, "id = ?", id).Error
}

func (r *RolesDao) Take(db *gorm.DB, where ...interface{}) *sysmdel.SysRoles {
	ret := &sysmdel.SysRoles{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *RolesDao) GetByName(db *gorm.DB, name string) *sysmdel.SysRoles {
	return r.Take(db, "name = ?", name)
}

func (r *RolesDao) GetById(db *gorm.DB, id int64) *sysmdel.SysRoles {
	return r.Take(db, "id = ?", id)
}
func (r *RolesDao) Create(db *gorm.DB, roles *sysmdel.SysRoles) (*sysmdel.SysRoles, error) {
	//表不存则创建表
	r.IsHasTable(db)
	if r.GetByName(db, roles.Name) != nil {
		return nil, errors.New("角色已存在")
	}
	err := db.Create(roles).Error
	return roles, err
}
func (r *RolesDao) IsHasTable(db *gorm.DB) {
	if db.Migrator().HasTable(&sysmdel.SysRoles{}) == false {
		db.Migrator().CreateTable(&sysmdel.SysRoles{})
	}
}

func (r *RolesDao) AnyRole(db *gorm.DB, roleId int64) (rolesList []sysmdel.SysRoles) {
	var SysRoles = &sysmdel.SysRoles{Id: roleId}
	db.Model(SysRoles).Where(SysRoles).Preload("Permissions").Find(&rolesList)
	return rolesList
}

func (r *RolesDao) RoleInfo(db *gorm.DB, roleId int64) (roles sysmdel.SysRoles) {
	var SysRoles = &sysmdel.SysRoles{Id: roleId}
	db.Model(SysRoles).Where(SysRoles).Preload("Permissions").First(&roles)
	return
}

func (r *RolesDao) RemoveRole(db *gorm.DB, roleId int64) error {
	if err := r.Delete(db, roleId); err == nil {
		roleErr := RolePermissionRepository.DeleteRolePermission(db, roleId)
		if roleErr == nil {
			return roleErr
		}
		return err
	}
	return nil
}
