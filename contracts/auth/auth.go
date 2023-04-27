package auth

import (
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/saas/sysmodel"
	"github.com/kataras/iris/v12"
)

//go:generate mockery --name=Auth
type Auth interface {
	Guard(name string) Auth
	Parse(ctx iris.Context) error
	User(ctx iris.Context) *sysmdel.PlatformAdmin
	Login(ctx iris.Context, UserLogin requests.UserLogin) (newAdmin *sysmdel.PlatformAdmin, token string, err error)
	LoginUsingID(ctx iris.Context, userId int64) (newAdmin *sysmdel.PlatformAdmin, token string, err error)
	Refresh(ctx iris.Context) (token string, err error)
	Logout(ctx iris.Context) bool
}
