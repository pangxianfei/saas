package requests

import "gitee.com/pangxianfei/frame/library/tmaic"

// ApplyFor 申请应用权限验证器
type ApplyFor struct {
	AppId int64 `json:"app_id" validate:"required,gt=0"`
}

func (r *ApplyFor) Messages() map[string][]string {
	messages := tmaic.M{
		"AppId": []string{
			"required:应用ID不为空",
			"gt:应用ID不合法",
		},
	}
	return messages
}
