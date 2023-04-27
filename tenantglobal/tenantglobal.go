package tenantglobal

import (
	"sync"

	"gitee.com/pangxianfei/saas/sysmodel"
)

var (
	InstanceDBMutex      sync.Mutex
	InstanceDBMapsObject = make(map[string]*sysmdel.InstanceObjectDB)
	InstanceDBMap        sync.Map
)
