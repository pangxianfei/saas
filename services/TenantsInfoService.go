package services

import (
	"errors"
	"time"

	"github.com/pangxianfei/saas/repositories"
	"github.com/pangxianfei/saas/requests"
	"github.com/pangxianfei/saas/simple"
	"github.com/pangxianfei/saas/sysmodel"
)

var TenantsInfoService = new(tenantsInfoService)

type tenantsInfoService struct {
}

func (s *tenantsInfoService) GetById(id int64) *sysmdel.PlatformAdmin {
	return repositories.AdminRepository.GetById(simple.DB(), id)
}

// GetByMobile 根据手机号查找
func (s *tenantsInfoService) GetByMobile(mobile string) *sysmdel.PlatformAdmin {
	return repositories.AdminRepository.GetByMobile(simple.DB(), mobile)
}

// GetByTenantName 根据户名查找
func (s *tenantsInfoService) GetByTenantName(TenantName string) *sysmdel.TenantsInfo {
	return repositories.TenantsRepository.GetByTenantName(TenantName)
}

// Create 存入DB
func (s *tenantsInfoService) Create(RegisterInfo requests.UserRegister) (Tenants *sysmdel.TenantsInfo, err error) {
	//保存至DB
	TenantsInfo := &sysmdel.TenantsInfo{
		TenantName: RegisterInfo.TenantName,
		Mobile:     RegisterInfo.Mobile,
		Status:     1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	Tenants, err = repositories.TenantsRepository.Create(TenantsInfo)
	if err != nil {
		return nil, errors.New("租户信息创建失败")
	}
	return Tenants, nil
}
