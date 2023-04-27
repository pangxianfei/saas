package services

import (
	"errors"

	"gitee.com/pangxianfei/saas/requests"

	"github.com/kataras/iris/v12"
	"github.com/pangxianfei/saas/consts"
	"github.com/pangxianfei/saas/simple"

	"github.com/pangxianfei/saas/buffer"
	"github.com/pangxianfei/saas/repositories"
	"github.com/pangxianfei/saas/sysmodel"
)

var UserService = new(userService)

type userService struct {
}

// GetByMobile 根据用户名查找
func (s *userService) GetByMobile(mobile string) *sysmdel.PlatformAdmin {
	return repositories.UserRepository.GetByMobile(simple.DB(), mobile)
}

// GetById 根据用户名查找
func (s *userService) GetById(id int64) *sysmdel.PlatformAdmin {
	return repositories.UserRepository.GetById(simple.DB(), id)
}

// SignIn 登录
func (s *userService) SignIn(ctx iris.Context, UserLogin requests.UserLogin) (*sysmdel.PlatformAdmin, string, error) {
	Admin := s.GetByMobile(UserLogin.Mobile)
	if Admin == nil || Admin.Id <= 0 {
		return nil, "", errors.New("用户不存在或被禁用")
	}
	//检验密码比较耗时 大约90毫秒
	if !simple.ValidatePassword(Admin.Password, UserLogin.Password) {
		return nil, "", errors.New("密码错误")
	}

	tokenSTR, loginErr := s.UnifyLogin(ctx, Admin)
	if loginErr != nil {
		return nil, "", loginErr
	}
	return Admin, tokenSTR, nil
}

// LoginUsingID 登录
func (s *userService) LoginUsingID(ctx iris.Context, adminId int64) (*sysmdel.PlatformAdmin, string, error) {
	Admin := s.GetById(adminId)
	if Admin == nil || Admin.Id <= 0 {
		return nil, "", errors.New("用户不存在或被禁用")
	}
	tokenSTR, loginErr := s.UnifyLogin(ctx, Admin)
	if loginErr != nil {
		return nil, "", loginErr
	}
	return Admin, tokenSTR, nil
}

// UnifyLogin 统一处理登录
func (s *userService) UnifyLogin(ctx iris.Context, admin *sysmdel.PlatformAdmin) (string, error) {
	if admin == nil || admin.Status == consts.StatusDisable || admin.Status == consts.StatusDeleted {

		return "", errors.New("用户不存在或被禁用")
	}

	tokenSTR, jwtErr := JwtService.InitMiddleware(admin)
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

	whereDb := &sysmdel.RetrieveDB{TenantsId: admin.TenantId, Status: 1}

	Result := simple.DB().Model(&sysmdel.InstanceDb{}).Where(whereDb).Find(&InstanceDB)

	if Result.RowsAffected <= 0 {
		return "", errors.New("租户应用不存在")
	}
	//缓存登陆用户信息
	buffer.UserCache.Set(admin)
	return tokenSTR, nil
}

func (s *userService) GetUserId(ctx iris.Context) int64 {
	if token, tErr := ctx.User().GetRaw(); tErr == nil {
		UserToken := token.(*sysmdel.Token)
		if UserToken.TenantId > 0 && UserToken.UserId > 0 {
			return UserToken.UserId
		}
	}
	return 0
}
func (s *userService) GetAdminInfo(ctx iris.Context) *sysmdel.PlatformAdmin {
	if token, tErr := ctx.User().GetRaw(); tErr == nil {
		UserToken := token.(*sysmdel.Token)
		if UserToken.TenantId > 0 && UserToken.UserId > 0 {
			return s.GetById(UserToken.UserId)
		}
	}
	return nil
}
