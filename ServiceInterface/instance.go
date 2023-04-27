package ServiceInterface

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type instance interface {
	GetTenantDb(tenantId int64, appId int64) *gorm.DB
	GetTenantUserDb(tenantId int64) (*gorm.DB, error)
	GetTenantId(ctx iris.Context) int64
}
