package sysmdel

import (
	"gitee.com/pangxianfei/frame/kernel/zone"
)

// SysRoles 角色表
type SysRoles struct {
	Id          int64             `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name        string            `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	TenantsId   int64             `gorm:"column:tenants_id;type:int(11) unsigned;NOT NULL" json:"tenants_id"`
	CreatedAt   zone.Time         `gorm:"column:created_at"`
	UpdatedAt   zone.Time         `gorm:"column:updated_at"`
	Permissions []RolePermissions `gorm:"foreignKey:RoleId"`
}

// TableName 指定表
func (s *SysRoles) TableName() string {
	return "sys_roles"
}
