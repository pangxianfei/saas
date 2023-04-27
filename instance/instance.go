package instance

import (
	"errors"
	"fmt"
	"strings"

	"gitee.com/pangxianfei/frame/kernel/debug"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"

	"gitee.com/pangxianfei/saas/tenant"

	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/frame/library/tmaic"
	"gitee.com/pangxianfei/frame/simple"
	"gitee.com/pangxianfei/frame/simple/date"

	"gitee.com/pangxianfei/saas/repositories"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/saas/saasapp"
	"gitee.com/pangxianfei/saas/services"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

type Instance struct{}

func (i *Instance) GetTenantAdminInfo(AdminId int64) *sysmdel.PlatformAdmin {
	return repositories.AdminRepository.GetById(simple.DB(), AdminId)
}

// CreateUser 保存租户信息
func (i *Instance) CreateUser(UserRegister requests.UserRegister, tenantsInfo *sysmdel.TenantsInfo) (Admin *sysmdel.PlatformAdmin, err error) {

	hasAdmin := services.RegisterService.GetByMobile(simple.DB(), UserRegister.Mobile)
	if hasAdmin != nil {
		return nil, errors.New("此手机号已被注册")
	}

	createAdmin := &sysmdel.PlatformAdmin{}
	TenantsErr := simple.DB().Transaction(func(txDb *gorm.DB) error {
		createAdmin.Nickname = tenantsInfo.Mobile
		createAdmin.TenantId = tenantsInfo.ID
		createAdmin.Mobile = tenantsInfo.Mobile
		createAdmin.UserName = tenantsInfo.Mobile
		createAdmin.Password = simple.EncodePassword(UserRegister.Password)
		createAdmin.Status = consts.StatusOk
		createAdmin.UserType = 2
		createAdmin.CreateTime = date.NowTimestamp()
		createAdmin.UpdateTime = date.NowTimestamp()
		if err := repositories.AdminRepository.Create(txDb, createAdmin); err != nil {
			return err
		}
		return nil
	})
	return createAdmin, TenantsErr
}

// CreateAppInstance 创建应用
func (i *Instance) CreateAppInstance(UserRegister requests.UserRegister, Admin *sysmdel.PlatformAdmin) (newInstanceDb []sysmdel.InstanceDb, err error) {

	//获取应用列表
	AppList := services.AppInfoService.GetByAppCreateList()

	if len(AppList) <= 0 {
		return nil, errors.New("租户：" + Admin.UserName + " 数据库出始化失败")
	}
	var InstanceDB sysmdel.InstanceDb

	tenantDbWhere := &sysmdel.RetrieveDB{
		Status: 1,
		Code:   "Create",
	}
	Result := simple.DB().Debug().Model(&sysmdel.InstanceDb{}).Where(tenantDbWhere).First(&InstanceDB)
	if Result.RowsAffected <= 0 {
		return nil, errors.New("租户： 数据库出始化失败")
	}

	dbCreateUser := fmt.Sprintf(consts.TENANTID_USER_DB_KEY, Admin.TenantId)

	CreateUserPass := fmt.Sprintf(consts.TENANTID_USER_DB_PASSWORD_KEY, Admin.TenantId, Admin.Mobile)

	dbCreatePassword := tmaic.MD5(CreateUserPass)

	debug.Dump("实例密码:" + dbCreatePassword)

	for _, appinfo := range AppList {
		saasDb := new(saasapp.App).SetTenantsDb(InstanceDB.TenantsId, InstanceDB.AppId)
		//创建应用数据库
		TenantDbName := fmt.Sprintf("tenant_db_%d_%s", Admin.TenantId, strings.ToLower(appinfo.Key))
		saasDb.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci", TenantDbName))

		//本地访问授权数据库给租户
		GrantLocalhost := fmt.Sprintf("GRANT ALL PRIVILEGES ON `%s`.* TO '%s'@'%s' WITH GRANT OPTION", TenantDbName, dbCreateUser, "%")
		saasDb.Exec(GrantLocalhost)
		//保存租户应用DB,创建应用数据库连接信息
		createInstanceDb := sysmdel.InstanceDb{
			TenantsId:          Admin.TenantId,
			Code:               appinfo.Key,
			AppName:            appinfo.Name,
			UseType:            "save",
			AppId:              appinfo.Id,
			Host:               InstanceDB.Host,
			DriverName:         InstanceDB.DriverName,
			Port:               InstanceDB.Port,
			Prefix:             InstanceDB.Prefix,
			DbName:             TenantDbName,
			Dbuser:             dbCreateUser,
			Charset:            InstanceDB.Charset,
			Collation:          InstanceDB.Collation,
			SetmaxIdleconns:    InstanceDB.SetmaxIdleconns,
			Setmaxopenconns:    InstanceDB.Setmaxopenconns,
			Setconnmaxlifetime: InstanceDB.Setconnmaxlifetime,
			Password:           dbCreatePassword,
			Intranet:           InstanceDB.Intranet,
			OuterNet:           InstanceDB.OuterNet,
			Status:             1,
		}
		checkInstanceDB := &sysmdel.InstanceDb{
			TenantsId: Admin.TenantId,
			AppId:     appinfo.Id,
		}
		result := simple.DB().Where(checkInstanceDB).First(checkInstanceDB)
		if result.RowsAffected <= 0 {
			simple.DB().Create(&createInstanceDb)
			newInstanceDb = append(newInstanceDb, createInstanceDb)
		}
	}
	return newInstanceDb, err
}
func (i *Instance) CreateDatabaseUserName(UserRegister requests.UserRegister, Admin *sysmdel.PlatformAdmin) error {

	tenantDbWhere := &sysmdel.RetrieveDB{Status: 1, Code: "Create"}

	var InstanceDB sysmdel.InstanceDb
	Result := simple.DB().Debug().Model(&sysmdel.InstanceDb{}).Where(tenantDbWhere).First(&InstanceDB)

	if Result.RowsAffected <= 0 {
		return errors.New("租户： 数据库出始化失败")
	}

	TenantsErr := simple.DB().Transaction(func(tx *gorm.DB) error {

		saasDb := new(saasapp.App).SetTenantsDb(InstanceDB.TenantsId, InstanceDB.AppId)

		dbCreateUser := fmt.Sprintf(consts.TENANTID_USER_DB_KEY, Admin.TenantId)

		CreateUserPass := fmt.Sprintf(consts.TENANTID_USER_DB_PASSWORD_KEY, Admin.TenantId, Admin.Mobile)

		dbCreatePassword := tmaic.MD5(CreateUserPass)
		//debug.Dump("实例密码:2" + dbCreatePassword)

		CreateUser := fmt.Sprintf(consts.MYSQL_CREATE_DB_USER_SQL, dbCreateUser, "%", dbCreatePassword)
		debug.Dd(CreateUser)
		saasDb.Exec(CreateUser)
		//内网访问权限
		//IntranetCreateUser := fmt.Sprintf("CREATE USER '%s'@'%s' IDENTIFIED BY '%s'", dbCreateUser, InstanceDB.Intranet, dbCreatePassword)
		//debug.Dd(IntranetCreateUser)
		//saasDb.Exec(IntranetCreateUser)

		//外网访问权限
		//OuterNetCreateUser := fmt.Sprintf("CREATE USER '%s'@'%s' IDENTIFIED BY '%s'", dbCreateUser, InstanceDB.OuterNet, dbCreatePassword)
		//debug.Dd(OuterNetCreateUser)
		//saasDb.Exec(OuterNetCreateUser)
		return nil
	})

	return TenantsErr

}

// CreateDBuser 后期部分 创建租户数据库帐号密码
func (i *Instance) CreateDBuser(AdminId int64) error {

	var Admin *sysmdel.PlatformAdmin
	var InstanceDB sysmdel.InstanceDb

	if Admin = repositories.AdminRepository.GetById(simple.DB(), AdminId); Admin == nil {
		return errors.New("管理用户不存在")
	}

	tenantDbWhere := &sysmdel.RetrieveDB{Status: 1, Code: "Create"}

	Result := simple.DB().Debug().Model(&sysmdel.InstanceDb{}).Where(tenantDbWhere).First(&InstanceDB)

	if Result.RowsAffected <= 0 {

		return errors.New("租户： 数据库出始化失败")

	}

	TenantsErr := simple.DB().Transaction(func(tx *gorm.DB) error {

		//连接数据库服务器
		saasDb := new(saasapp.App).SetTenantsDb(InstanceDB.TenantsId, InstanceDB.AppId)

		dbCreateUser := fmt.Sprintf(consts.TENANTID_USER_DB_KEY, Admin.TenantId)

		CreateUserPass := fmt.Sprintf(consts.TENANTID_USER_DB_PASSWORD_KEY, Admin.TenantId, Admin.Mobile)

		dbCreatePassword := tmaic.MD5(CreateUserPass)

		CreateUser := fmt.Sprintf(consts.MYSQL_CREATE_DB_USER_SQL, dbCreateUser, "%", dbCreatePassword)

		saasDb.Exec(CreateUser)
		//内网访问权限
		IntranetCreateUser := fmt.Sprintf(consts.MYSQL_CREATE_DB_USER_SQL, dbCreateUser, InstanceDB.Intranet, dbCreatePassword)

		saasDb.Exec(IntranetCreateUser)
		//外网访问权限
		if InstanceDB.Intranet != InstanceDB.OuterNet {
			OuterNetCreateUser := fmt.Sprintf(consts.MYSQL_CREATE_DB_USER_SQL, dbCreateUser, InstanceDB.OuterNet, dbCreatePassword)

			saasDb.Exec(OuterNetCreateUser)
		}

		return nil
	})

	return TenantsErr

}
func (i *Instance) CreateTenantsDatabase(AdminId int64) (newInstanceDb []sysmdel.InstanceDb, err error) {

	var Admin *sysmdel.PlatformAdmin

	if Admin = repositories.AdminRepository.GetById(simple.DB(), AdminId); Admin == nil {
		return nil, errors.New("管理用户不存在")
	}
	//获取应用列表
	AppList := services.AppInfoService.GetByAppCreateList()
	if len(AppList) <= 0 {
		return nil, errors.New("租户：" + Admin.UserName + " 数据库出始化失败")
	}

	var InstanceDB sysmdel.InstanceDb

	tenantDbWhere := &sysmdel.RetrieveDB{Status: 1, AppId: 1000001}

	Result := simple.DB().Model(&sysmdel.InstanceDb{}).Where(tenantDbWhere).First(&InstanceDB)

	if Result.RowsAffected <= 0 {

		return nil, errors.New("租户： 数据库出始化失败")
	}

	dbCreateUser := fmt.Sprintf(consts.TENANTID_USER_DB_KEY, Admin.TenantId)

	CreateUserPass := fmt.Sprintf(consts.TENANTID_USER_DB_PASSWORD_KEY, Admin.TenantId, Admin.Mobile)

	dbCreatePassword := tmaic.MD5(CreateUserPass)

	for _, appinfo := range AppList {

		saasDb := new(saasapp.App).SetTenantsDb(InstanceDB.TenantsId, InstanceDB.AppId)
		//创建应用数据库
		TenantDbName := fmt.Sprintf(consts.MYSQL_APP_DB_NAME_KEY, Admin.TenantId, strings.ToLower(appinfo.Key))

		saasDb.Exec(fmt.Sprintf(consts.MYSQL_CREATE_DATABASE_SQL, TenantDbName))

		//本地访问授权数据库给租户
		GrantLocalhost := fmt.Sprintf(consts.MYSQL_GRANT_USER_DB_SQL, TenantDbName, dbCreateUser, "%")

		saasDb.Exec(GrantLocalhost)
		//保存租户应用DB,创建应用数据库连接信息
		createInstanceDb := sysmdel.InstanceDb{
			TenantsId:          Admin.TenantId,
			Code:               appinfo.Key,
			AppName:            appinfo.Name,
			UseType:            "save",
			AppId:              appinfo.Id,
			Host:               InstanceDB.Host,
			DriverName:         InstanceDB.DriverName,
			Port:               InstanceDB.Port,
			Prefix:             InstanceDB.Prefix,
			DbName:             TenantDbName,
			Dbuser:             dbCreateUser,
			Charset:            InstanceDB.Charset,
			Collation:          InstanceDB.Collation,
			SetmaxIdleconns:    InstanceDB.SetmaxIdleconns,
			Setmaxopenconns:    InstanceDB.Setmaxopenconns,
			Setconnmaxlifetime: InstanceDB.Setconnmaxlifetime,
			Password:           dbCreatePassword,
			Intranet:           InstanceDB.Intranet,
			OuterNet:           InstanceDB.OuterNet,
			Status:             1,
		}
		checkInstanceDB := &sysmdel.InstanceDb{
			TenantsId: Admin.TenantId,
			AppId:     appinfo.Id,
		}
		result := simple.DB().Where(checkInstanceDB).First(checkInstanceDB)
		if result.RowsAffected <= 0 {
			simple.DB().Create(&createInstanceDb)
			newInstanceDb = append(newInstanceDb, createInstanceDb)
		}
	}
	return newInstanceDb, err

}

// CreateLoginAccount 创建登陆帐号
func (i *Instance) CreateLoginAccount(cxt iris.Context, UserName string, Mobile string, Password string) (*sysmdel.PlatformAdmin, error) {
	hasAdmin := services.RegisterService.GetByMobile(simple.DB(), Mobile)
	if hasAdmin != nil {
		return nil, errors.New("此手机号已被注册")
	}
	AdminInfo := services.UserTokenService.GetUserInfo(cxt)
	if AdminInfo == nil || AdminInfo.Id <= 0 {
		return nil, errors.New("非法请求")
	}
	createAdmin := &sysmdel.PlatformAdmin{}
	TenantsErr := simple.DB().Transaction(func(txDb *gorm.DB) error {
		createAdmin.TenantId = AdminInfo.TenantId
		createAdmin.Mobile = Mobile
		createAdmin.UserName = UserName
		createAdmin.Password = simple.EncodePassword(Password)
		createAdmin.Status = consts.StatusOk
		createAdmin.UserType = AdminInfo.UserType
		createAdmin.CreateTime = date.NowTimestamp()
		createAdmin.UpdateTime = date.NowTimestamp()
		if err := repositories.AdminRepository.Create(txDb, createAdmin); err != nil {
			return err
		}
		return nil
	})
	if TenantsErr != nil {
		return nil, errors.New("创建失败,请重试")
	}
	var Tenant tenant.Tenant
	_, err := Tenant.SynTenantUser(createAdmin)
	if err != nil {
		return nil, errors.New("创建失败,请重试")
	}
	return createAdmin, nil
}
