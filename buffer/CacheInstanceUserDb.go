package buffer

import (
	"encoding/json"
	"fmt"

	"gitee.com/pangxianfei/frame/kernel/cache"
	"gitee.com/pangxianfei/frame/kernel/debug"
	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/frame/simple"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

var CacheInstanceUserDb = new(CacheInstanceUser)

type CacheInstanceUser struct{}

// GetCacheInstanceUserDb 获取缓存数据
func (c *CacheInstanceUser) GetCacheInstanceUserDb(TenantsId int64) (InstanceDb *sysmdel.InstanceDb) {
	TenantIdUserDbAppKey := fmt.Sprintf(consts.TENANTID_USER_DB_APP_KEY, TenantsId, consts.UserDb)
	TenantIdUserDbAppKey = simple.MD5(TenantIdUserDbAppKey)
	instanceUserDbData := cache.GetString(TenantIdUserDbAppKey)
	if instanceUserDbData == "" {
		return nil
	}
	err := simple.InterfaceToStruct(instanceUserDbData, &InstanceDb)
	if err != nil {
		return nil
	}
	return InstanceDb
}

// SetCacheInstanceUserDb 设置缓存数据
func (c *CacheInstanceUser) SetCacheInstanceUserDb(instanceDb *sysmdel.InstanceDb) bool {
	instanceUserDbData, _ := json.Marshal(instanceDb)
	TenantIdUserDbAppKey := fmt.Sprintf(consts.TENANTID_USER_DB_APP_KEY, instanceDb.TenantsId, consts.UserDb)
	TenantIdUserDbAppKey = simple.MD5(TenantIdUserDbAppKey)
	debug.Dd(TenantIdUserDbAppKey)
	return cache.AddTokenCache(TenantIdUserDbAppKey, instanceUserDbData)
}
