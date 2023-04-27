package repositories

import (
	"gorm.io/gorm"

	"github.com/pangxianfei/saas/sysmodel"
)

var AdminRepository = new(adminRepository)

type adminRepository struct {
}

func (r *adminRepository) Take(db *gorm.DB, where ...interface{}) *sysmdel.PlatformAdmin {
	ret := &sysmdel.PlatformAdmin{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *adminRepository) GetById(db *gorm.DB, AdminId int64) *sysmdel.PlatformAdmin {
	return r.Take(db, "id = ?", AdminId)
}

func (r *adminRepository) GetByMobile(db *gorm.DB, mobile string) *sysmdel.PlatformAdmin {
	return r.Take(db, "mobile = ?", mobile)
}

func (r *adminRepository) Create(db *gorm.DB, admin *sysmdel.PlatformAdmin) (err error) {
	err = db.Create(admin).Error
	return
}

func (r *adminRepository) TenantUserRegister(db *gorm.DB, admin *sysmdel.PlatformAdmin) (*sysmdel.TenantAdmin, error) {

	TenantUser := &sysmdel.TenantAdmin{
		TenantId:         admin.TenantId,
		Mobile:           admin.Mobile,
		UserName:         admin.UserName,
		Email:            admin.Email,
		EmailVerified:    admin.EmailVerified,
		Nickname:         admin.Nickname,
		Avatar:           admin.Avatar,
		BackgroundImage:  admin.BackgroundImage,
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
			return nil, err
		}
	}

	err := db.Create(TenantUser).Error
	return TenantUser, err
}
