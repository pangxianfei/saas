package ServiceInterface

import (
	"gitee.com/pangxianfei/saas/requests"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

type appInfo interface {
	GetByName(mobile string) *sysmdel.AppInfo
	GetByList() []sysmdel.AppInfo
	Create(appInfo requests.AppInfo) (AppInfo *sysmdel.AppInfo, err error)
	GetByAppCreateList() []sysmdel.AppInfo
	GetStartApplication() []sysmdel.AppInfo
}
