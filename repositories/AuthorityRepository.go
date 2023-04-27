package repositories

import (
	"gorm.io/gorm"

	"gitee.com/pangxianfei/saas/sysmodel"
)

var AuthorityRepository = new(authorityDao)

type authorityDao struct {
}

func (r *authorityDao) Take(db *gorm.DB, where ...interface{}) *sysmdel.Authority {
	ret := &sysmdel.Authority{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *authorityDao) GetByName(db *gorm.DB, name string) *sysmdel.Authority {
	return r.Take(db, "name = ?", name)
}

// GetByAppIdList 列表
func (r *authorityDao) GetByAppIdList(db *gorm.DB, appId int64) []sysmdel.Authority {
	var AppList []sysmdel.Authority
	db.Where(&sysmdel.Authority{AppId: appId}).Find(&AppList)
	return AppList
}

func (r *authorityDao) Create(db *gorm.DB, AppInfo *sysmdel.Authority) (appInfo *sysmdel.Authority, err error) {
	r.IsHasTable(db)
	err = db.Create(AppInfo).Error
	return
}

func (r *authorityDao) IsHasTable(db *gorm.DB) {
	if db.Migrator().HasTable(&sysmdel.Authority{}) == false {
		_ = db.Migrator().CreateTable(&sysmdel.Authority{})
	}
}
