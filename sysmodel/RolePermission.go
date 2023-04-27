package sysmdel

// RolePermissions 角色拥有权限
type RolePermissions struct {
	Id           int64 `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	PermissionId int64 `gorm:"column:permission_id;type:int(11);index:idx_permission_id" json:"permission_id"`
	RoleId       int64 `gorm:"column:role_id;type:int(11);NOT NULL" json:"role_id"`
	TenantsId    int64 `gorm:"column:tenants_id;type:int(11);NOT NULL" json:"tenants_id"`
	AppId        int64 `gorm:"column:app_id;type:int(11);NOT NULL" json:"app_id"`
}

// TableName 指定表
func (s *RolePermissions) TableName() string {
	return "sys_role_has_permissions"
}
