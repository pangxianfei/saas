package sysmdel

type UserToken struct {
	Id         int64  `gorm:"primaryKey;autoIncrement;" json:"id" form:"id"`
	TenantId   int64  `gorm:"size:11;not null;<-:create" json:"TenantId" form:"TenantId"`
	Mobile     string `gorm:"size:11;;embeddedPrefix:my_" json:"mobile" form:"mobile"` // 手机
	Token      string `gorm:"type:text;" json:"token" form:"token"`
	UserId     int64  `gorm:"not null;index:idx_user_token_user_id;" json:"userId" form:"userId"`
	ExpiredAt  int64  `gorm:"not null;" json:"expiredAt" form:"expiredAt"`
	Status     int    `gorm:"not null;" json:"status" form:"status"`
	CreateTime int64  `gorm:"not null" json:"createTime" form:"createTime"`
	Md5Token   string `gorm:"size:255;unique;" json:"md5_token" form:"md5_token"` // token MD5
}

func (this *UserToken) TableName() string {
	return "sys_user_token"
}
