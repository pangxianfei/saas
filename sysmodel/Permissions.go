package sysmdel

import (
	"gitee.com/pangxianfei/frame/kernel/zone"
)

// Permissions 权限明细表
type Permissions struct {
	Id           int64  `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	PermissionId int64  `gorm:"column:permission_id;type:int(11);NOT NULL;uniqueIndex" json:"permission_id"`
	Pid          int64  `gorm:"type:int(11);not null;size:11" json:"pid" form:"pid"`
	AppId        int64  `gorm:"type:int(11);not null;size:11" json:"app_id" form:"app_id"`
	Name         string `gorm:"type:varchar(255);not null;" json:"name" form:"name"`
	MenuName     string `gorm:"type:varchar(255);" json:"menu_name" form:"menu_name"`
	IsMenu       int64  `gorm:"type:int(11);" json:"is_menu" form:"is_menu"`
	Description  string `gorm:"type:varchar(255);not null;" json:"description" form:"description"`
	//RegisterFileName  string        `gorm:"type:varchar(255);not null;" json:"register_file_name" form:"register_file_name"`
	//MainHandlerName   string        `gorm:"type:varchar(255);not null;" json:"main_handler_name" form:"main_handler_name"`
	Method        string `gorm:"type:varchar(255);not null;" json:"method" form:"method"`
	FormattedPath string `gorm:"type:varchar(255);not null;" json:"formatted_path" form:"formatted_path"`
	StaticPath    string `gorm:"type:varchar(255);not null;" json:"static_path" form:"static_path"`
	Path          string `gorm:"type:varchar(255);not null;" json:"path" form:"path"`
	//SourceFileName    string        `gorm:"type:varchar(255);not null;" json:"source_file_name" form:"source_file_name"`
	RouteName         string        `gorm:"type:varchar(255);not null;" json:"route_name" form:"route_name"`
	Status            int64         `gorm:"type:int(1);not null;default:0;" json:"status" form:"status"`
	Md5Value          string        `gorm:"type:varchar(255);not null;" json:"md5_value" form:"md5_value"`
	RouteNameMd5Value string        `gorm:"type:varchar(255);not null;" json:"route_name_md5_value" form:"route_name_md5_value"`
	Children          []Permissions `gorm:"foreignKey:Pid;references:PermissionId"  json:"children"`
	CreatedAt         zone.Time     `gorm:"column:created_at"`
	UpdatedAt         zone.Time     `gorm:"column:updated_at"`
}

// TableName 指定表
func (s *Permissions) TableName() string {
	return "sys_permissions"
}
