package services

import (
	"errors"

	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/frame/simple"
	"gitee.com/pangxianfei/saas/buffer"
	"gitee.com/pangxianfei/saas/saasapp"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var InstanceService = new(instanceService)

type instanceService struct {
}

func (i *instanceService) GetTenantDb(tenantId int64, appId int64) *gorm.DB {
	return new(saasapp.App).SetTenantsDb(tenantId, appId)
}
func (i *instanceService) GetTenantUserDb(tenantId int64) (*gorm.DB, error) {
	//获取用户系统应用ID
	var InstanceDB *sysmdel.InstanceDb
	newInstanceUserDB := buffer.CacheInstanceUserDb.GetCacheInstanceUserDb(tenantId)
	if newInstanceUserDB != nil && newInstanceUserDB.TenantsId == tenantId {
		InstanceDB = newInstanceUserDB
	} else {
		tenantDbWhere := &sysmdel.RetrieveDB{
			TenantsId: tenantId,
			Status:    1,
			Code:      consts.UserDb,
		}
		Result := simple.DB().Model(&sysmdel.InstanceDb{}).Where(tenantDbWhere).First(&InstanceDB)

		if Result.RowsAffected <= 0 {
			return nil, errors.New("租户数据库不存在")
		}
		//缓存用户系统应用 DB 信息
		buffer.CacheInstanceUserDb.SetCacheInstanceUserDb(InstanceDB)
	}

	db := new(saasapp.App).SetTenantsDb(tenantId, InstanceDB.AppId)
	return db, nil
}
func (i *instanceService) GetTenantId(ctx iris.Context) int64 {
	token, tErr := ctx.User().GetRaw()
	if tErr != nil {
		return 0
	}
	UserToken := token.(*sysmdel.Token)
	if UserToken.TenantId > 0 {
		return UserToken.TenantId
	}
	return 0
}
