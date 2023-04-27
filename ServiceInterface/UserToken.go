package ServiceInterface

import (
	"gitee.com/pangxianfei/saas/sysmodel"
	"github.com/kataras/iris/v12"
)

type UserToken interface {
	GetUserInfo(ctx iris.Context) (adminInfo *sysmdel.PlatformAdmin)
	Create(Admin *sysmdel.PlatformAdmin, token string) (*sysmdel.UserToken, error)
	Disable(ctx iris.Context) bool
	Parse(ctx iris.Context) error
}
