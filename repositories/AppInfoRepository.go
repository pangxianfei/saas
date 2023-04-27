package repositories

import (
	"gorm.io/gorm"

	"gitee.com/pangxianfei/saas/sysmodel"
)

var AppInfoRepository = new(appInfoRepository)

type appInfoRepository struct {
}

func (r *appInfoRepository) Take(db *gorm.DB, where ...interface{}) *sysmdel.AppInfo {
	ret := &sysmdel.AppInfo{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}
func (r *appInfoRepository) GetByName(db *gorm.DB, name string) *sysmdel.AppInfo {
	return r.Take(db, "name = ?", name)
}

func (r *appInfoRepository) Create(db *gorm.DB, AppInfo *sysmdel.AppInfo) (appInfo *sysmdel.AppInfo, err error) {
	r.IsHasTable(db)
	err = db.Create(AppInfo).Error
	return
}

func (r *appInfoRepository) GetByList(db *gorm.DB) []sysmdel.AppInfo {
	var AppList []sysmdel.AppInfo
	db.Find(&AppList)
	return AppList
}

func (r *appInfoRepository) GetByAppCreateList(db *gorm.DB) []sysmdel.AppInfo {
	var AppList []sysmdel.AppInfo
	db.Debug().Where("is_created = ?", 1).Find(&AppList)
	return AppList
}

func (r *appInfoRepository) GetStartApplication(db *gorm.DB) []sysmdel.AppInfo {
	var AppList []sysmdel.AppInfo
	db.Where("status = ?", 1).Find(&AppList)
	return AppList
}

func (r *appInfoRepository) IsHasTable(db *gorm.DB) {
	if db.Migrator().HasTable(&sysmdel.Permissions{}) == false {
		_ = db.Migrator().CreateTable(&sysmdel.Permissions{})
	}
}
