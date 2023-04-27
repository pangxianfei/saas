package buffer

import (
	"fmt"

	"gitee.com/pangxianfei/frame/kernel/cache"
	"gitee.com/pangxianfei/frame/kernel/debug"
	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/frame/simple"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

var CachePermission = new(Permission)

type Permission struct {
}

// GetCachePermission 获取缓存数据
func (c *Permission) GetCachePermission(UserId int64, RouteName string) bool {
	AdminPermissionKey := fmt.Sprintf(consts.ADMIN_PERMISSION_KEY, UserId, RouteName)
	AdminPermissionKey = simple.MD5(AdminPermissionKey)
	cacheData := cache.GetString(AdminPermissionKey)
	if cacheData == "" {
		return false
	}

	if len(cacheData) > 0 {
		return true
	}
	return false
}

// SetCachePermission 获取缓存数据
func (c *Permission) SetCachePermission(adminPermissions *sysmdel.AdminPermissions) bool {
	AdminPermissionKey := fmt.Sprintf(consts.ADMIN_PERMISSION_KEY, adminPermissions.UserId, adminPermissions.RouteName)
	AdminPermissionKey = simple.MD5(AdminPermissionKey)
	debug.Dd(AdminPermissionKey)
	return cache.AddTokenCache(AdminPermissionKey, adminPermissions.RouteName)
}
