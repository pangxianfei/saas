package ServiceInterface

import (
	"github.com/kataras/iris/v12"

	"gitee.com/pangxianfei/saas/requests"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

type user interface {
	GetByMobile(mobile string) *sysmdel.PlatformAdmin
	GetById(id int64) *sysmdel.PlatformAdmin
	SignIn(ctx iris.Context, UserLogin requests.UserLogin) (*sysmdel.PlatformAdmin, string, error)
	LoginUsingID(ctx iris.Context, adminId int64) (*sysmdel.PlatformAdmin, string, error)
	UnifyLogin(ctx iris.Context, admin *sysmdel.PlatformAdmin) (string, error)
	GetUserId(ctx iris.Context) int64
	GetAdminInfo(ctx iris.Context) *sysmdel.PlatformAdmin
}
