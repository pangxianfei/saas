package buffer

import (
	"encoding/json"
	"fmt"

	"gitee.com/pangxianfei/frame/kernel/cache"
	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/frame/simple"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

var CacheInstanceDb = new(CacheInstance)

type CacheInstance struct {
}

// GetCacheInstanceDb 获取缓存数据
func (c *CacheInstance) GetCacheInstanceDb(TenantsId int64, AppId int64) (InstanceDb *sysmdel.InstanceDb) {
	TenantIdDbAppKey := fmt.Sprintf(consts.TENANTID_DB_APP_KEY, TenantsId, AppId)
	TenantIdDbAppKey = simple.MD5(TenantIdDbAppKey)
	cacheData := cache.GetString(TenantIdDbAppKey)
	if cacheData == "" {
		return nil
	}
	if err := simple.InterfaceToStruct(cacheData, &InstanceDb); err != nil {
		return nil
	}
	if InstanceDb.ID > 0 && InstanceDb.TenantsId > 0 && InstanceDb.AppId > 0 {
		return InstanceDb
	}
	return nil
}

// SetCacheInstanceDb 设置缓存数据
func (c *CacheInstance) SetCacheInstanceDb(instanceDb *sysmdel.InstanceDb) bool {
	if instanceDb == nil || instanceDb.ID <= 0 {
		return false
	}
	instanceDbData, _ := json.Marshal(instanceDb)
	TenantIdDbAppKey := fmt.Sprintf(consts.TENANTID_DB_APP_KEY, instanceDb.TenantsId, instanceDb.AppId)
	TenantIdDbAppKey = simple.MD5(TenantIdDbAppKey)
	return cache.AddTokenCache(TenantIdDbAppKey, instanceDbData)
}
