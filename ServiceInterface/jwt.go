package ServiceInterface

import sysmdel "gitee.com/pangxianfei/saas/sysmodel"

type jwt interface {
	InitMiddleware(Admin *sysmdel.PlatformAdmin) (string, error)
}
