package repositories

import (
	"github.com/pangxianfei/saas/sysmodel"
	"gorm.io/gorm"
)

var UserRepository = new(userRepository)

type userRepository struct {
}

func (r *userRepository) Take(db *gorm.DB, where ...interface{}) *sysmdel.PlatformAdmin {
	ret := &sysmdel.PlatformAdmin{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userRepository) GetByMobile(db *gorm.DB, mobile string) *sysmdel.PlatformAdmin {
	return r.Take(db, "mobile = ?", mobile)
}

func (r *userRepository) GetById(db *gorm.DB, id int64) *sysmdel.PlatformAdmin {
	return r.Take(db, "id = ?", id)
}
