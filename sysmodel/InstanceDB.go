package sysmdel

import (
	"gitee.com/pangxianfei/frame/kernel/zone"
)

// InstanceDb 租户数据信息模型
type InstanceDb struct {
	ID                 int64     `gorm:"column:id;primary_key;auto_increment"`
	TenantsId          int64     `gorm:"column:tenants_id;type:int unsigned;not null"`
	Code               string    `gorm:"column:code;type:varchar(255);not null"`
	AppName            string    `gorm:"column:app_name;type:varchar(255);not null"`
	UseType            string    `gorm:"column:use_type;type:varchar(255);not null"`
	AppId              int64     `gorm:"column:app_id;type:int unsigned;not null"`
	Host               string    `gorm:"column:host;type:varchar(255);not null"`
	DriverName         string    `gorm:"column:drivername;type:varchar(255);not null"`
	Port               int       `gorm:"column:port;type:int unsigned;not null"`
	Prefix             string    `gorm:"column:prefix;type:varchar(255);not null"`
	DbName             string    `gorm:"column:dbname;type:varchar(255);not null"`
	Dbuser             string    `gorm:"column:dbuser;type:varchar(255);not null"`
	Charset            string    `gorm:"column:charset;type:varchar(255);not null"`
	Collation          string    `gorm:"column:collation;type:varchar(255);not null"`
	SetmaxIdleconns    int       `gorm:"column:setmaxIdleconns;type:int unsigned;not null"`
	Setmaxopenconns    int       `gorm:"column:setmaxopenconns;type:int unsigned;not null"`
	Setconnmaxlifetime int       `gorm:"column:setconnmaxlifetime;type:int unsigned;not null"`
	Status             int       `gorm:"column:status;type:int(1);"`
	NetSwitch          int       `gorm:"column:net_switch;type:int(1);"`
	Password           string    `gorm:"column:password;type:varchar(255);not null"`
	Intranet           string    `gorm:"column:intranet;type:varchar(255);not null"`  //内网访问地址
	OuterNet           string    `gorm:"column:outer_net;type:varchar(255);not null"` //外网访问地址
	CreatedAt          zone.Time `gorm:"column:created_at"`
	UpdatedAt          zone.Time `gorm:"column:updated_at"`
}

// TableName 指定表
func (instanceDb *InstanceDb) TableName() string {
	return "sys_instance_db"
}
