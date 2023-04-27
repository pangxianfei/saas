package tenant

import (
	"github.com/kataras/iris/v12"

	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/saas/sysmodel"
)

type Tenant interface {
	SynTenantUser(platformAdmin *sysmdel.PlatformAdmin) (*sysmdel.TenantAdmin, error)
	Add(cxt iris.Context, UserRegister requests.UserRegister) (Tenants *sysmdel.TenantsInfo, err error)
	GetByTenantName(TenantName string) *sysmdel.TenantsInfo
	GetTenantAdminInfo(id int64) *sysmdel.PlatformAdmin
}
