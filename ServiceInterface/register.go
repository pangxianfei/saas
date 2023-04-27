package ServiceInterface

import (
	"gitee.com/pangxianfei/saas/sysmodel"
	"gorm.io/gorm"
)

type register interface {
	GetByMobile(db *gorm.DB, mobile string) *sysmdel.PlatformAdmin
	GetBySubMobile(db *gorm.DB, mobile string) *sysmdel.TenantAdmin
	SynTenantUser(admin *sysmdel.PlatformAdmin) (*sysmdel.TenantAdmin, error)
}
