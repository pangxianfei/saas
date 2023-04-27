package middleware

import (
	"strings"

	"gitee.com/pangxianfei/frame/console"
	"gitee.com/pangxianfei/frame/library/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"

	"gitee.com/pangxianfei/saas/buffer"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/saas/response"
	"gitee.com/pangxianfei/saas/saasapp"
	"gitee.com/pangxianfei/saas/sysmodel"
)

func TenantTenantMiddleware(ctx iris.Context, AppId int64) context.Handler {
	var UserAppHeader requests.AppHeaderAuthorization
	if err := ctx.ReadHeaders(&UserAppHeader); err != nil {
		_ = ctx.JSON(response.ErrorTokenInvalidation())
		return nil
	}
	TokenModel := jwt.Get(ctx).(*sysmdel.Token)
	standardClaims := jwt.GetVerifiedToken(ctx).StandardClaims
	timeLeft := standardClaims.Timeleft()
	if TokenModel.TenantId <= 0 || TokenModel.UserId <= 0 {
		_ = ctx.JSON(response.ErrorUnauthorized())
		return nil
	}
	//获取当前登陆的用户 判断TOKEN是否有效
	token := strings.ReplaceAll(UserAppHeader.Authorization, "Bearer ", "")
	if buffer.UserTokenCache.IsTokenInvalid(token) {
		_ = ctx.JSON(response.ErrorTokenInvalidation())
		return nil
	}
	//if TokenModel.UserId != 1 {
	sassDb := new(saasapp.App).SetTenantsDb(TokenModel.TenantId, AppId)
	if sassDb == nil && TokenModel.UserId != 1 {
		_ = ctx.JSON(response.ErrorUnregisteredTenantAppDb())
		return nil
	}
	//}
	if config.Instance.TenantLog {
		console.Println(console.CODE_WARNING, " "+console.Sprintf(console.CODE_WARNING, "应用编号: %d %s  %s %s", AppId, ctx.Method(), ctx.Path(), timeLeft))
	}
	ctx.Values().Set("TenantId", TokenModel.TenantId)
	ctx.Values().Set("AppId", AppId)
	ctx.Values().Set("UserId", TokenModel.UserId)
	ctx.Next()
	return nil
}
