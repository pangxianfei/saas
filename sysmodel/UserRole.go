package sysmdel

type UserRoles struct {
	Id        int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleId    int64 `gorm:"column:role_id;type:int(11) UNSIGNED" json:"roleId"`
	UserId    int64 `gorm:"column:user_id;type:int(11)" json:"userId"`
	TenantsId int64 `gorm:"column:tenants_id;type:int(11) UNSIGNED" json:"tenantsId"`
}

func (u *UserRoles) TableName() string {
	return "sys_user_has_roles"
}
