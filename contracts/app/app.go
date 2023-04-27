package saasapp

import (
	"gitee.com/pangxianfei/saas/sysmodel"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type SaasApp interface {
	Initiation(ctx iris.Context) (TenantDB *gorm.DB)
	SetTenantsDb(TenantsId int64, AppId int64) (TenantDB *gorm.DB)
	SetSourceData(CurrInstanceDB *sysmdel.InstanceDb) (TenantDB *gorm.DB)
	MysqlOpenDb(Instance *sysmdel.InstanceDb) (tenant *gorm.DB)
	OutConsole(title string, serial int64, TenantsId int64, AppId int64, DbName string)
	//Refresh(ctx iris.Context) (token string, err error)
	//Logout(ctx iris.Context) error
}
