package auth

import (
	contractAuth "gitee.com/pangxianfei/saas/contracts/auth"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/saas/services"
	"gitee.com/pangxianfei/saas/sysmodel"
	"github.com/kataras/iris/v12"
)

type Auth struct {
}

func NewAuth(guard string) *Auth {
	return &Auth{}
}

func (app *Auth) Guard(name string) contractAuth.Auth {
	return NewAuth(name)
}

func (app *Auth) Parse(ctx iris.Context) error {
	return services.UserTokenService.Parse(ctx)
}

func (app *Auth) User(ctx iris.Context) *sysmdel.PlatformAdmin {

	return services.UserTokenService.GetUserInfo(ctx)
}

func (app *Auth) Login(ctx iris.Context, UserLogin requests.UserLogin) (newAdmin *sysmdel.PlatformAdmin, token string, err error) {
	return services.UserService.SignIn(ctx, UserLogin)
}

func (app *Auth) LoginUsingID(ctx iris.Context, userId int64) (newAdmin *sysmdel.PlatformAdmin, token string, err error) {
	return services.UserService.LoginUsingID(ctx, userId)
}

func (app *Auth) Refresh(ctx iris.Context) (token string, err error) {
	return "", nil
}

func (app *Auth) Logout(ctx iris.Context) bool {
	return services.UserTokenService.Disable(ctx)
}
