package sysmdel

type AdminPermissions struct {
	Id                int64  `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	UserId            int64  `gorm:"column:user_id;type:int" json:"userId"`
	PermissionId      int64  `gorm:"column:permission_id;type:int" json:"permissionId"`
	TenantsId         int64  `gorm:"column:tenants_id;type:int" json:"tenantsId"`
	AppId             int64  `gorm:"column:app_id;type:int" json:"appId"`
	RoleId            int64  `gorm:"column:role_id;type:int(11);NOT NULL" json:"role_id"`
	RouteNameMd5Value string `gorm:"type:varchar(255);not null;" json:"route_name_md5_value" form:"route_name_md5_value"`
	RouteName         string `gorm:"type:varchar(255);not null;" json:"route_name" form:"route_name"`
}

func (s *AdminPermissions) TableName() string {
	return "sys_admin_has_permissions"
}
