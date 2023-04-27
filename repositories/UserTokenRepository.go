package repositories

import (
	"gitee.com/pangxianfei/frame/simple"
	"gitee.com/pangxianfei/frame/simple/sqlcmd"
	"gorm.io/gorm"

	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

var UserTokenRepository = new(userTokenRepository)

type userTokenRepository struct {
}

func (r *userTokenRepository) GetByToken(db *gorm.DB, token string) *sysmdel.UserToken {
	if len(token) == 0 {
		return nil
	}
	return r.Take(db, "token = ?", token)
}

func (r *userTokenRepository) Get(db *gorm.DB, id int64) *sysmdel.UserToken {
	ret := &sysmdel.UserToken{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) Take(db *gorm.DB, where ...interface{}) *sysmdel.UserToken {
	ret := &sysmdel.UserToken{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) Find(db *gorm.DB, cnd *sqlcmd.Cnd) (list []sysmdel.UserToken) {
	cnd.Find(db, &list)
	return
}

func (r *userTokenRepository) FindOne(db *gorm.DB, cnd *sqlcmd.Cnd) *sysmdel.UserToken {
	ret := &sysmdel.UserToken{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []sysmdel.UserToken, paging *sqlcmd.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userTokenRepository) FindPageByCnd(db *gorm.DB, cnd *sqlcmd.Cnd) (list []sysmdel.UserToken, paging *sqlcmd.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &sysmdel.UserToken{})

	paging = &sqlcmd.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userTokenRepository) Create(db *gorm.DB, t *sysmdel.UserToken) (err error) {
	r.IsHasTable(db)
	err = db.Debug().Create(t).Error
	return
}

func (r *userTokenRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&sysmdel.UserToken{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userTokenRepository) UpdateColumnToken(db *gorm.DB, tokenKey string, name string, value interface{}) (err error) {
	err = db.Debug().Model(&sysmdel.UserToken{}).Where("md5_token = ?", tokenKey).UpdateColumn(name, value).Error
	return
}

func (r *userTokenRepository) IsHasTable(db *gorm.DB) {
	if db.Migrator().HasTable(&sysmdel.UserToken{}) == false {
		_ = db.Migrator().CreateTable(&sysmdel.UserToken{})
	}
}
