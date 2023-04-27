package ServiceInterface

import (
	"gitee.com/pangxianfei/saas/requests"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

type tenant interface {
	GetByMobile(mobile string) *sysmdel.PlatformAdmin
	GetByTenantName(TenantName string) *sysmdel.TenantsInfo
	Create(RegisterInfo requests.UserRegister) (Tenants *sysmdel.TenantsInfo, err error)
}
