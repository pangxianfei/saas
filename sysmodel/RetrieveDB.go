package sysmdel

// RetrieveDB 检索租户对象
type RetrieveDB struct {
	AppId     int64 // 应用id
	TenantsId int64 // 租户id
	UseType   string
	Status    int64  // 状态
	Code      string // 状态
}
