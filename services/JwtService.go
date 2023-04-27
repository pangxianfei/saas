package services

import (
	"time"

	"gitee.com/pangxianfei/frame/library/config"

	"github.com/pangxianfei/saas/sysmodel"

	ginJwt "gitee.com/pangxianfei/frame/utils/jwt"
	irisJwt "github.com/kataras/iris/v12/middleware/jwt"
)

var JwtService = new(jwtService)

type jwtService struct {
}

func (t *jwtService) InitMiddleware(Admin *sysmdel.PlatformAdmin) (string, error) {
	authSignKey := []byte(config.GetString("auth.sign_key"))
	var tokenTime time.Duration
	tokenTime = time.Duration(config.GetInt64("cache.token_time"))
	signer := irisJwt.NewSigner(irisJwt.HS256, authSignKey, tokenTime*time.Minute)
	claims := sysmdel.Token{
		UserId:   Admin.Id,
		Mobile:   Admin.Mobile,
		TenantId: Admin.TenantId,
		Username: Admin.Mobile,
		Email:    Admin.Mobile,
		Iss:      "tmaic",
	}
	tokenSTR, err := signer.Sign(claims)
	if err != nil {
		return "", err
	}
	return string(tokenSTR), nil
}

func (t *jwtService) InitGinMiddleware(Admin *sysmdel.PlatformAdmin) (string, error) {

	newJwt := ginJwt.NewJWT()

	token, err := newJwt.CreateToken(Admin.Id, Admin.Mobile, Admin.TenantId)
	if err != nil {
		return "", err
	}
	return token, nil
}
