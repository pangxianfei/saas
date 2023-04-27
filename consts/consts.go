package consts

const (
	DefaultTokenExpireDays = 7   // 用户登录token默认有效期
	SummaryLen             = 256 // 摘要长度
)

// 系统配置
const (
	SysConfigSiteTitle          = "siteTitle"          // 站点标题
	SysConfigSiteDescription    = "siteDescription"    // 站点描述
	SysConfigSiteKeywords       = "siteKeywords"       // 站点关键字
	SysConfigSiteNavs           = "siteNavs"           // 站点导航
	SysConfigSiteNotification   = "siteNotification"   // 站点公告
	SysConfigRecommendTags      = "recommendTags"      // 推荐标签
	SysConfigUrlRedirect        = "urlRedirect"        // 是否开启链接跳转
	SysConfigScoreConfig        = "scoreConfig"        // 分数配置
	SysConfigDefaultNodeId      = "defaultNodeId"      // 发帖默认节点
	SysConfigArticlePending     = "articlePending"     // 是否开启文章审核
	SysConfigTopicCaptcha       = "topicCaptcha"       // 是否开启发帖验证码
	SysConfigUserObserveSeconds = "userObserveSeconds" // 新用户观察期
	SysConfigTokenExpireDays    = "tokenExpireDays"    // 登录Token有效天数
	SysConfigLoginMethod        = "loginMethod"        // 登录方式
)

// EntityType
const (
	EntityArticle = "article"
	EntityTopic   = "topic"
	EntityComment = "comment"
	EntityTweet   = "tweet"
	EntityUser    = "user"
	EntityCheckIn = "checkIn"
)

// 用户角色
const (
	RoleOwner = "owner" // 站长
	RoleAdmin = "admin" // 管理员
	RoleUser  = "user"  // 用户
)

// 用户角色id
const (
	RoleAdministrator int64 = 1 // 超级管理员
)

// 操作类型
const (
	OpTypeCreate          = "create"
	OpTypeDelete          = "delete"
	OpTypeUpdate          = "update"
	OpTypeForbidden       = "forbidden"
	OpTypeRemoveForbidden = "removeForbidden"
)

// 状态
const (
	StatusOk      = 1    // 正常
	StatusDisable = 1000 // 禁用
	StatusDeleted = 1001 // 删除
	StatusPending = 3    // 待审核
)

// 用户类型
const (
	UserTypeNormal = 0        // 普通用户
	UserTypeGzh    = 1        // 公众号用户
	UserDb         = "User"   // 租户下用户db库标志
	createDb       = "create" // 平台使用
	createUserDb   = "Sys"
)

// 内容类型
const (
	ContentTypeHtml     = "html"
	ContentTypeMarkdown = "markdown"
	ContentTypeText     = "text"
)

// 消息状态
const (
	MsgStatusUnread   = 1 // 消息未读
	MsgStatusHaveRead = 2 // 消息已读
)

// 应用实例
const APP_DB_KEY = "%d"
const APP_DB_CACHE_KEY = "instance_db_%d"

// 字符串 permission InstanceDb
const TENANTID_DB = "%d_%d"
const USER_CACAHE_KEY = "admin_cache_%d"
const ADMIN_PERMISSION_KEY = "admin_permission_cache_%d_%s"
const TENANTID_DB_APP_KEY = "instance_db_%d_%d"
const TENANTID_USER_DB_APP_KEY = "instance_user_db_%d_%s"

const MYSQL_CREATE_DB_USER_SQL = "CREATE USER `%s`@`%s` IDENTIFIED WITH mysql_native_password BY '%s';"
const MYSQL_GRANT_USER_DB_SQL = "GRANT ALTER,ALTER Routine,CREATE,CREATE Routine,CREATE TEMPORARY TABLES,CREATE VIEW,DELETE,DROP,EVENT,EXECUTE,GRANT OPTION,INDEX,INSERT,LOCK TABLES,REFERENCES,SELECT,SHOW VIEW,TRIGGER,UPDATE ON `%s`.* TO `%s` @`%s`;"
const MYSQL_CREATE_DATABASE_SQL = "CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci"

const TENANTID_USER_DB_KEY = "tenant_db_%d"          //数据用户名
const MYSQL_APP_DB_NAME_KEY = "tenant_db_%d_%s"      //应用数据库
const TENANTID_USER_DB_PASSWORD_KEY = "tenant_%d_%s" //数据库密码格式
