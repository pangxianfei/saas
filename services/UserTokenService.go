package services

import (
	"errors"
	"fmt"
	"time"

	"gitee.com/pangxianfei/frame/library/config"
	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/frame/library/tmaic"
	"gitee.com/pangxianfei/frame/simple"
	"gitee.com/pangxianfei/frame/simple/date"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"

	"gitee.com/pangxianfei/saas/buffer"
	"gitee.com/pangxianfei/saas/repositories"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

var UserTokenService = new(userTokenService)

type userTokenService struct {
}

// GetUserInfo 获取当前登录用户的id
func (s *userTokenService) GetUserInfo(ctx iris.Context) (adminInfo *sysmdel.PlatformAdmin) {
	token := s.GetUserToken(ctx)
	//获取登陆token缓存
	userToken := buffer.UserTokenCache.Get(token)
	if userToken == nil {
		return nil
	}

	userInfoKey := fmt.Sprintf(consts.USER_CACAHE_KEY, userToken.UserId, userToken.TenantId)
	adminInfo = buffer.UserCache.GetCacheKey(userInfoKey)
	if adminInfo.Id <= 0 {
		adminInfo = UserService.GetById(userToken.UserId)
	}
	// 没找到授权
	if userToken == nil || userToken.Status == consts.StatusDeleted {
		return nil
	}
	// 授权过期
	if userToken.ExpiredAt >= date.NowTimestamp() {
		return nil
	}
	return adminInfo
}

// GetUserToken 从请求体中获取UserToken
func (s *userTokenService) GetUserToken(ctx iris.Context) string {
	userToken := ctx.GetHeader("Authorization")
	if len(userToken) > 0 {
		return userToken[7:]
	}
	return ""
}

// Create 存入DB
func (s *userTokenService) Create(Admin *sysmdel.PlatformAdmin, token string) (*sysmdel.UserToken, error) {
	var iat = time.Now().Unix()
	var exp = config.GetInt64("cache.token_time")
	//保存至DB
	UserToken := &sysmdel.UserToken{
		Token:      token,
		UserId:     Admin.Id,
		TenantId:   Admin.TenantId,
		Mobile:     Admin.Mobile,
		ExpiredAt:  iat + exp,
		Status:     0,
		CreateTime: iat,
		Md5Token:   tmaic.MD5(token),
	}
	err := repositories.UserTokenRepository.Create(simple.DB(), UserToken)
	if err != nil {
		return nil, errors.New("token存入DB失败")
	}
	buffer.UserTokenCache.Invalidate(token)
	return UserToken, nil
}

// Disable 禁用
func (s *userTokenService) Disable(ctx iris.Context) bool {
	token := s.GetUserToken(ctx)
	//设置token缓存失效
	_, tokenKey := buffer.UserTokenCache.Invalidate(token)
	statusErr := repositories.UserTokenRepository.UpdateColumnToken(simple.DB(), tokenKey, "status", consts.StatusDeleted)
	return statusErr == nil
}

// Parse 解释token
func (s *userTokenService) Parse(ctx iris.Context) error {
	s.GetUserToken(ctx)
	TokenModel := jwt.Get(ctx).(*sysmdel.Token)
	TenantId := ctx.Values().Get("TenantId").(int64)

	if TokenModel.UserId > 0 && TokenModel.TenantId > 0 && TokenModel.TenantId == TenantId {
		return nil
	}
	return errors.New("cache support is required")
}
