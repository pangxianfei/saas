package tenant

import (
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/pangxianfei/saas/repositories"
	"github.com/pangxianfei/saas/requests"
	"github.com/pangxianfei/saas/services"
	"github.com/pangxianfei/saas/simple"
	"github.com/pangxianfei/saas/sysmodel"
	"time"
)

type Tenant struct {
}

// SynTenantUser 同步帐号、密码至 用户应用db
func (app *Tenant) SynTenantUser(platformAdmin *sysmdel.PlatformAdmin) (*sysmdel.TenantAdmin, error) {
	return services.RegisterService.SynTenantUser(platformAdmin)
}

func (app *Tenant) GetById(id int64) *sysmdel.PlatformAdmin {
	return services.TenantsInfoService.GetById(id)
}

func (app *Tenant) GetByMobile(mobile string) *sysmdel.PlatformAdmin {
	return services.TenantsInfoService.GetByMobile(mobile)
}

// GetByTenantName 根据名称查询
func (app *Tenant) GetByTenantName(TenantName string) *sysmdel.TenantsInfo {
	return services.TenantsInfoService.GetByTenantName(TenantName)
}

// Add 创建租户信息
func (app *Tenant) Add(cxt iris.Context, UserRegister requests.UserRegister) (Tenants *sysmdel.TenantsInfo, err error) {
	//查询租户名称是否占用
	TenantsInfo := services.TenantsInfoService.GetByTenantName(UserRegister.TenantName)

	if TenantsInfo != nil {
		return nil, errors.New("租户名称：" + UserRegister.TenantName + " 已被占用")
	}

	// 验证手机号
	if simple.IsMobile(UserRegister.Mobile) && services.TenantsInfoService.GetByMobile(UserRegister.Mobile) != nil {
		return nil, errors.New("手机号：" + UserRegister.Mobile + " 已被占用")
	}

	//保存至DB
	newTenantsInfo := &sysmdel.TenantsInfo{
		TenantName: UserRegister.TenantName,
		Mobile:     UserRegister.Mobile,
		Status:     1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	//写入DB
	Tenants, err = repositories.TenantsRepository.Create(newTenantsInfo)
	if err != nil {
		return nil, errors.New("租户信息创建失败")
	}

	return Tenants, nil
}

// GetTenantAdminInfo 平台租户登陆帐号信息
func (app *Tenant) GetTenantAdminInfo(id int64) *sysmdel.PlatformAdmin {
	return services.UserService.GetById(id)
}
