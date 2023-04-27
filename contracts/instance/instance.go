package instance

import (
	"github.com/kataras/iris/v12"

	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/saas/sysmodel"
)

type Instance interface {
	GetTenantAdminInfo(AdminId int64) *sysmdel.PlatformAdmin
	CreateUser(UserRegister requests.UserRegister, tenantsInfo *sysmdel.TenantsInfo) (Admin *sysmdel.PlatformAdmin, err error)
	CreateAppInstance(UserRegister requests.UserRegister, Admin *sysmdel.PlatformAdmin) (newInstanceDb []sysmdel.InstanceDb, err error)
	CreateDatabaseUserName(UserRegister requests.UserRegister, Admin *sysmdel.PlatformAdmin) error
	CreateDBuser(AdminId int64) error
	CreateTenantsDatabase(AdminId int64) (newInstanceDb []sysmdel.InstanceDb, err error)
	CreateLoginAccount(cxt iris.Context, UserName string, Mobile string, Password string) (*sysmdel.PlatformAdmin, error)
}
