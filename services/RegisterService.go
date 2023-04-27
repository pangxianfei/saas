package services

import (
	"errors"

	"github.com/pangxianfei/saas/consts"
	"github.com/pangxianfei/saas/simple"
	"gorm.io/gorm"

	"github.com/pangxianfei/saas/repositories"
	"github.com/pangxianfei/saas/saasapp"
	"github.com/pangxianfei/saas/sysmodel"
)

var RegisterService = new(registerService)

type registerService struct {
}

// GetByMobile 根据用户名查找
func (s *registerService) GetByMobile(db *gorm.DB, mobile string) *sysmdel.PlatformAdmin {
	return repositories.AdminRepository.GetByMobile(db, mobile)
}

// GetBySubMobile 根据用户名查找
func (s *registerService) GetBySubMobile(db *gorm.DB, mobile string) *sysmdel.TenantAdmin {
	return repositories.TenantAdminRepository.GetByMobile(db, mobile)
}

// SynTenantUser 同步帐号
func (s *registerService) SynTenantUser(admin *sysmdel.PlatformAdmin) (*sysmdel.TenantAdmin, error) {
	//获取用户系统应用ID
	var InstanceDB sysmdel.InstanceDb
	tenantDbWhere := &sysmdel.RetrieveDB{
		TenantsId: admin.TenantId,
		Status:    1,
		Code:      consts.UserDb,
	}
	Result := simple.DB().Model(&sysmdel.InstanceDb{}).Where(tenantDbWhere).First(&InstanceDB)
	if Result.RowsAffected > 0 {
		db := new(saasapp.App).SetTenantsDb(admin.TenantId, InstanceDB.AppId)
		TenantUser := s.GetBySubMobile(db, admin.Mobile)
		if TenantUser != nil {
			return nil, errors.New("帐号已存,无需同步")
		}
		newTenantUser, err := repositories.AdminRepository.TenantUserRegister(db, admin)
		if err == nil {
			return newTenantUser, nil
		}
	}
	return nil, errors.New("同步失败")
}
