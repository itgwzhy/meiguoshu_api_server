package model

import (
	"github.com/MEIGUOSHU/api-server/packages/model/base"
)

//咨询
type Consultation struct {
	base.Model
	Name    string	//称呼
	City    string	//城市
	Company string	//公司
	Mobile  string	//手机号码
	Email   string	//邮箱
	Content string	//咨询内容
}
