package requests

import "gitee.com/pangxianfei/frame/library/tmaic"

// UserLogin 登陆验证器
type UserLogin struct {
	//TenantId int64  `json:"tenantId" validate:"required,gt=0"`
	Mobile   string `json:"mobile" validate:"required,len=11"`
	Password string `json:"password" validate:"required,min=7,max=24"`
}

func (r *UserLogin) Messages() map[string][]string {
	messages := tmaic.M{
		"Mobile": []string{
			"required:手机号为必填项",
			"len:手机号必须必须是11位的号码 gt:",
		},
		"Password": []string{
			"required:密码为必填项",
			"min:密码必须大于 min:",
			"max:密码必须小于 max:",
		},
	}
	return messages
}
