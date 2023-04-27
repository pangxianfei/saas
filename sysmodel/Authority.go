package sysmdel

// Authority 权限模型
type Authority struct {
	Id                int64  `gorm:"column:id;type:int(11) UNSIGNED"`
	Pid               int64  `gorm:"column:pid;type:int(11)"`
	AppId             int64  `gorm:"column:app_id;type:int(11)"`
	Description       string `gorm:"column:description;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	RegisterFileName  string `gorm:"column:register_file_name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	MenuName          string `gorm:"column:menu_name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	IsMenu            int64  `gorm:"column:is_menu;type:int(11)"` // 权限类型
	Name              string `gorm:"column:name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	MainHandlerName   string `gorm:"column:main_handler_name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	Method            string `gorm:"column:method;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	FormattedPath     string `gorm:"column:formatted_path;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	StaticPath        string `gorm:"column:static_path;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	Path              string `gorm:"column:path;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	SourceFileName    string `gorm:"column:source_file_name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	RouteName         string `gorm:"column:route_name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	RouteNameMd5Value string `gorm:"column:route_name_md5_value;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	Status            int64  `gorm:"column:status;type:int(11)"`
	Md5Value          string `gorm:"column:md5_value;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	UpdatedAt         string `gorm:"column:updated_at;type:timestamp"`
	CreatedAt         string `gorm:"column:created_at;type:timestamp"`
}

// TableName 指定表
func (Auth *Authority) TableName() string {
	return "public_permissions"
}
