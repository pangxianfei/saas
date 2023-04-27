package tenantglobal

import (
	"database/sql"
	"fmt"
	"gitee.com/pangxianfei/frame/database"

	"gitee.com/pangxianfei/frame/library/config"
	"gitee.com/pangxianfei/frame/library/console"
	"gitee.com/pangxianfei/frame/library/consts"
	"gitee.com/pangxianfei/frame/simple"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"

	"gitee.com/pangxianfei/saas/buffer"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
)

func Initiation(ctx iris.Context) (TenantDB *gorm.DB) {
	TenantsId := ctx.Values().Get("TenantId").(int64)
	AppId := ctx.Values().Get("AppId").(int64)
	return SetTenantsDb(TenantsId, AppId)
}

func SetTenantsDb(TenantsId int64, AppId int64) (TenantDB *gorm.DB) {
	if TenantsId <= 0 || AppId <= 0 {
		return nil
	}
	var CurrInstanceDB *sysmdel.InstanceDb

	var TenantsAppDBNo string = fmt.Sprintf(consts.TENANTID_DB, TenantsId, AppId)

	newTenantIdAppDB := buffer.CacheInstanceDb.GetCacheInstanceDb(TenantsId, AppId)
	if newTenantIdAppDB != nil && newTenantIdAppDB.ID > 0 && newTenantIdAppDB.TenantsId == TenantsId && newTenantIdAppDB.AppId == AppId {
		CurrInstanceDB = newTenantIdAppDB
	} else {
		tenantDbWhere := &sysmdel.RetrieveDB{TenantsId: TenantsId, AppId: AppId}
		if err := simple.DB().Debug().Where(tenantDbWhere).Find(&CurrInstanceDB).Error; err != nil {
			return nil
		}
		buffer.CacheInstanceDb.SetCacheInstanceDb(CurrInstanceDB)
	}

	UserTenantsDB, hasDbConn := InstanceDBMap.Load(TenantsAppDBNo)
	if hasDbConn {
		var ThemTenantsDB *sysmdel.InstanceObjectDB = UserTenantsDB.(*sysmdel.InstanceObjectDB)
		if ThemTenantsDB != nil && ThemTenantsDB.TenantsId == TenantsId && ThemTenantsDB.AppId == AppId && CurrInstanceDB.TenantsId == TenantsId && CurrInstanceDB.AppId == AppId {
			OutConsole("缓存", 1, ThemTenantsDB.TenantsId, ThemTenantsDB.AppId, ThemTenantsDB.DbName)
			TenantDB = ThemTenantsDB.Db
		}
	}
	if TenantDB != nil {
		return TenantDB.Session(&gorm.Session{})
	}

	OutConsole("出始化", 3, CurrInstanceDB.TenantsId, CurrInstanceDB.AppId, CurrInstanceDB.DbName)
	return SetSourceData(CurrInstanceDB)
}

func SetSourceData(CurrInstanceDB *sysmdel.InstanceDb) (InstanceDB *gorm.DB) {
	var mysqlStr string
	var dber *sql.DB
	TenantsAppDBNo := fmt.Sprintf(consts.TENANTID_DB, CurrInstanceDB.TenantsId, CurrInstanceDB.AppId)
	InstanceDB = MysqlOpenDb(CurrInstanceDB)

	SaveConnectionObject := sysmdel.InstanceObjectDB{}
	SaveConnectionObject.DbName = CurrInstanceDB.DbName
	SaveConnectionObject.AppName = CurrInstanceDB.AppName
	SaveConnectionObject.TenantsId = CurrInstanceDB.TenantsId
	SaveConnectionObject.AppId = CurrInstanceDB.AppId
	SaveConnectionObject.Prefix = CurrInstanceDB.Prefix
	SaveConnectionObject.DriverName = CurrInstanceDB.DriverName
	SaveConnectionObject.Db = InstanceDB
	SaveConnectionObject.Dber = dber
	SaveConnectionObject.ConnStr = mysqlStr
	InstanceDBMutex.Lock()
	//defer InstanceDBMutex.Unlock()
	//InstanceDBMapsObject[TenantsAppDBNo] = &SaveConnectionObject
	InstanceDBMap.Store(TenantsAppDBNo, &SaveConnectionObject)
	return InstanceDB
}

func MysqlOpenDb(Instance *sysmdel.InstanceDb) (TenantDB *gorm.DB) {
	var DbDns string
	//如果是使用外网IP
	if Instance.NetSwitch == 2 {
		DbDns = Instance.OuterNet
	} else {
		//内网IP
		DbDns = Instance.Host
	}
	_, TenantDB = database.OpenDB(Instance.DriverName, DbDns, Instance.Dbuser, Instance.Password, Instance.DbName, Instance.Port, Instance.Prefix, Instance.Charset, Instance.SetmaxIdleconns, Instance.Setmaxopenconns, Instance.Setconnmaxlifetime)
	return TenantDB
}
func OutConsole(title string, serial int64, TenantsId int64, AppId int64, DbName string) {
	var consoleSTR string
	consoleSTR = " " + console.Sprintf(console.CODE_WARNING, "序号%d:连接类型:%s 租户:%d - 应用ID:%d  - 数据库:%s", serial, title, TenantsId, AppId, DbName)
	if config.Instance.TenantDb {
		console.Println(console.CODE_WARNING, consoleSTR)
	}
}
