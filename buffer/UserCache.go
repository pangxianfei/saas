package buffer

import (
	"encoding/json"
	"fmt"

	"gitee.com/pangxianfei/frame/kernel/cache"
	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/frame/simple"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

type userCache struct {
}

var UserCache = newUserCache()

func newUserCache() *userCache {
	return &userCache{}
}

func (c *userCache) Get(Admin *sysmdel.PlatformAdmin) *sysmdel.PlatformAdmin {
	userInfoKey := fmt.Sprintf(consts.USER_CACAHE_KEY, Admin.Id, Admin.TenantId)

	userInfoKey = simple.MD5(userInfoKey)

	adminInfo := cache.Get(userInfoKey, nil)

	return adminInfo.(*sysmdel.PlatformAdmin)
}

func (c *userCache) GetCacheKey(userInfoKey string) *sysmdel.PlatformAdmin {

	if len(userInfoKey) <= 0 {
		return nil
	}
	userInfoKey = simple.MD5(userInfoKey)
	adminInfoValue := cache.Get(userInfoKey)
	var adminInfo *sysmdel.PlatformAdmin

	newData := simple.InterfaceToString(adminInfoValue)

	err := json.Unmarshal([]byte(newData), &adminInfo)

	if err != nil {
		return nil
	}
	return adminInfo
}

func (c *userCache) Set(Admin *sysmdel.PlatformAdmin) bool {

	if Admin.Id <= 0 {
		return false
	}

	userData, _ := json.Marshal(Admin)

	userInfoKey := fmt.Sprintf(consts.USER_CACAHE_KEY, Admin.Id, Admin.TenantId)

	tokenKey := simple.MD5(userInfoKey)

	if cache.AddTokenCache(tokenKey, userData) {
		return true
	}
	return false
}

func (c *userCache) Invalidate(userId int64) {

}
