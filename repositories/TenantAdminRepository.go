package repositories

import (
	"gorm.io/gorm"

	"github.com/pangxianfei/saas/sysmodel"
)

var TenantAdminRepository = new(tenantUserRepository)

type tenantUserRepository struct {
}

func (r *tenantUserRepository) Take(db *gorm.DB, where ...interface{}) *sysmdel.TenantAdmin {
	ret := &sysmdel.TenantAdmin{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *tenantUserRepository) GetByMobile(db *gorm.DB, mobile string) *sysmdel.TenantAdmin {
	return r.Take(db, "mobile = ?", mobile)
}

func (r *tenantUserRepository) Create(db *gorm.DB, admin *sysmdel.TenantAdmin) (err error) {
	//表不存则创建表
	r.IsHasTable(db)
	return db.Create(admin).Error
}

func (r *tenantUserRepository) TenantUserRegister(db *gorm.DB, admin *sysmdel.TenantAdmin) (err error) {

	TenantUser := &sysmdel.TenantAdmin{
		TenantId:        admin.TenantId,
		Mobile:          admin.Mobile,
		UserName:        admin.UserName,
		Email:           admin.Email,
		EmailVerified:   admin.EmailVerified,
		Nickname:        admin.Nickname,
		Avatar:          admin.Avatar,
		BackgroundImage: admin.BackgroundImage,
		//Password:         admin.Password,
		HomePage:         admin.HomePage,
		Description:      admin.Description,
		Score:            admin.Score,
		Status:           admin.Status,
		TopicCount:       admin.TopicCount,
		CommentCount:     admin.CommentCount,
		Roles:            admin.Roles,
		UserType:         admin.UserType,
		ForbiddenEndTime: admin.ForbiddenEndTime,
		CreateTime:       admin.CreateTime,
		UpdateTime:       admin.UpdateTime,
	}

	var createStatus bool = false
	createStatus = db.Debug().Migrator().HasTable(&sysmdel.TenantAdmin{})
	if createStatus == false {
		err := db.Migrator().CreateTable(&sysmdel.TenantAdmin{})
		if err != nil {
			return err
		}
	}

	err = db.Create(TenantUser).Error
	return
}
func (r *tenantUserRepository) IsHasTable(db *gorm.DB) {
	if db.Migrator().HasTable(&sysmdel.Permissions{}) == false {
		_ = db.Migrator().CreateTable(&sysmdel.Permissions{})
	}
}
