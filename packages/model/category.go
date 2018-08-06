package model

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/packages/model/base"
)

//分类
type Category struct {
	base.Model
	name  string //名称
	tty   int64  //类型 1：产品中心分类 2：研发推广分类
	pid   int64  //父iD
	level int64  //公司等级
}
