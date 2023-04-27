package sysmdel

type Token struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   int64  `json:"UserId"`
	TenantId int64  `json:"TenantId"`
	AppId    int64  `json:"AppId"`
	Mobile   string `json:"Mobile"`
	UserType int    `json:"type"`
	Iss      string `json:"iss"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
}
