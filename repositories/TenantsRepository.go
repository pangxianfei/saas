package repositories

import (
	"github.com/pangxianfei/saas/simple"
	"github.com/pangxianfei/saas/sysmodel"
	"gorm.io/gorm"
)

var TenantsRepository = new(tenantsRepository)

type tenantsRepository struct {
}

func (r *tenantsRepository) Take(db *gorm.DB, where ...interface{}) *sysmdel.TenantsInfo {
	ret := &sysmdel.TenantsInfo{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *tenantsRepository) GetByMobile(mobile string) *sysmdel.TenantsInfo {
	return r.Take(simple.DB(), "mobile = ?", mobile)
}

func (r *tenantsRepository) GetByTenantName(TenantName string) *sysmdel.TenantsInfo {
	return r.Take(simple.DB(), "tenant_name = ?", TenantName)
}

func (r *tenantsRepository) Create(TenantsInfo *sysmdel.TenantsInfo) (Tenants *sysmdel.TenantsInfo, err error) {
	var createStatus bool = false
	db := simple.DB()
	createStatus = db.Migrator().HasTable(&sysmdel.TenantsInfo{})
	if createStatus == false {
		err := db.Migrator().CreateTable(&sysmdel.TenantsInfo{})
		if err != nil {
			return nil, err
		}
	}
	err = db.Create(TenantsInfo).Error
	return TenantsInfo, err
}
