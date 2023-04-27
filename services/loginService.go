package services

import (
	"errors"

	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/frame/simple"

	"gitee.com/pangxianfei/frame/request"

	"gitee.com/pangxianfei/saas/buffer"
	"gitee.com/pangxianfei/saas/repositories"
	"gitee.com/pangxianfei/saas/sysmodel"
)

var LoginService = new(loginService)

type loginService struct {
}

// GetByMobile 根据用户名查找
func (login *loginService) GetByMobile(mobile string) *sysmdel.PlatformAdmin {
	return repositories.UserRepository.GetByMobile(simple.DB(), mobile)
}

// GetById 根据用户名查找
func (login *loginService) GetById(id int64) *sysmdel.PlatformAdmin {
	return repositories.UserRepository.GetById(simple.DB(), id)
}

func (login *loginService) User(ctx request.Context) *sysmdel.PlatformAdmin {
	return nil
}

func (login *loginService) Login(ctx request.Context, Mobile string, Password string) (newAdmin *sysmdel.PlatformAdmin, token string, err error) {
	Admin := login.GetByMobile(Mobile)

	if Admin == nil || Admin.Id <= 0 {

		return nil, "", errors.New("用户不存在")
	}
	//检验密码比较耗时 大约90毫秒
	if !simple.ValidatePassword(Admin.Password, Password) {

		return nil, "", errors.New("密码错误")
	}

	tokenSTR, loginErr := login.UnifyLogin(Admin)

	if loginErr != nil {

		return nil, "", loginErr
	}

	return Admin, tokenSTR, nil
}

func (login *loginService) LoginUsingID(ctx request.Context, userId int64) (newAdmin *sysmdel.PlatformAdmin, token string, err error) {

	return nil, "", nil
}

// UnifyLogin  基于gin JWT 统一处理登录
func (login *loginService) UnifyLogin(admin *sysmdel.PlatformAdmin) (UserTokenSTR string, err error) {
	if admin == nil || admin.Status == consts.StatusDisable || admin.Status == consts.StatusDeleted {

		return "", errors.New("用户不存在或被禁用")
	}
	tokenSTR, jwtErr := JwtService.InitGinMiddleware(admin)

	if jwtErr != nil {

		return "", jwtErr
	}
	//写入DB
	UserToken, dbErr := UserTokenService.Create(admin, tokenSTR)
	if dbErr != nil {

		return "", dbErr

	}
	//缓存token
	buffer.UserTokenCache.SetCacheUserToken(tokenSTR, UserToken)
	//出始化租户 DB 连接对象
	var InstanceDB []sysmdel.InstanceDb

	whereDb := &sysmdel.RetrieveDB{Status: 1}

	Result := simple.DB().Model(&sysmdel.InstanceDb{}).Where(whereDb).Find(&InstanceDB)

	if Result.RowsAffected <= 0 {

		return "", errors.New("应用不存在")

	}
	//缓存登陆用户信息
	buffer.UserCache.Set(admin)
	return tokenSTR, err
}
