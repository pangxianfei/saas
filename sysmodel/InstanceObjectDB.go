package sysmdel

import (
	"database/sql"
	"gorm.io/gorm"
)

type InstanceObjectDB struct {
	DbName     string   // 数据库连接名称
	AppName    string   //应用名
	TenantsId  int64    // 租户id
	AppId      int64    // 应用id
	ConnStr    string   // 连接串
	DriverName string   // 驱动名称
	Db         *gorm.DB // 数据库连接对象
	Dber       *sql.DB  // 数据库连接对象
	Prefix     string   //表前缀
}
