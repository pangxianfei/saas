package buffer

import (
	"encoding/json"

	"gitee.com/pangxianfei/frame/kernel/cache"
	"gitee.com/pangxianfei/frame/simple"

	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

var UserTokenCache = newUserTokenCache()

type userTokenCache struct {
}

func newUserTokenCache() *userTokenCache {
	return &userTokenCache{}
}

// Get 获取缓存数据
func (c *userTokenCache) Get(token string) (UserToken *sysmdel.UserToken) {

	if len(token) == 0 {
		return nil
	}
	tokenKey := simple.MD5(token)

	//读取缓存
	cacheData := cache.GetString(tokenKey)

	if len(cacheData) <= 0 {
		return nil
	}
	if err := simple.InterfaceToStruct(cacheData, &UserToken); err != nil {
		return nil
	}
	return
}

func (c *userTokenCache) SetCacheUserToken(token string, userToken *sysmdel.UserToken) {
	userTokenData, _ := json.Marshal(userToken)
	tokenKey := simple.MD5(token)
	cache.AddTokenCache(tokenKey, userTokenData)
}

func (c *userTokenCache) Invalidate(token string) (Forget bool, tokenKey string) {
	if len(token) <= 0 {
		return false, ""
	}
	tokenKey = simple.MD5(token)
	return cache.Forget(tokenKey), tokenKey
}

func (c *userTokenCache) IsTokenInvalid(token string) bool {
	if len(token) <= 0 {
		return false
	}
	tokenKey := simple.MD5(token)
	tokenValue := cache.Get(tokenKey)
	if tokenValue == nil || len(tokenValue.(string)) <= 0 {
		return true
	}
	return false
}
