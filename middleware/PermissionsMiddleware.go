package middleware

import (
	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/response"
	"gitee.com/pangxianfei/saas/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var Permissions context.Handler

func GetPermissions() {
	Permissions = func(ctx iris.Context) {
		UserId := services.UserService.GetUserId(ctx)
		if UserId == consts.RoleAdministrator {
			//超级管理员放行
		} else if paas.Gate.HasPermission(ctx, ctx.GetCurrentRoute().Name()) == false {
			ctx.StatusCode(iris.StatusForbidden)
			_ = ctx.JSON(response.ErrorNoHaveAuthority())
			return
		}
		ctx.Next()
	}
}
