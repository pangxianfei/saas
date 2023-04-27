package middleware

import (
	"gitee.com/pangxianfei/frame/library/config"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"

	"gitee.com/pangxianfei/saas/response"
	"gitee.com/pangxianfei/saas/sysmodel"
)

// LoginMiddleware 检查登陆验证器
func LoginMiddleware() context.Handler {

	verifier := jwt.NewVerifier(jwt.HS256, []byte(config.GetString("auth.sign_key")))
	verifier.WithDefaultBlocklist()
	verifier.ErrorHandler = func(ctx *context.Context, err error) {
		ctx.StatusCode(401)
		_ = ctx.JSON(response.ErrorTokenInvalidation())
		return
	}
	return verifier.Verify(func() interface{} {
		return new(sysmdel.Token)
	})
}
