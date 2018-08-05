package model

import (
	"github.com/MEIGUOSHU/web-api/packages/model/base"
)

// 品牌资讯
type News struct {
	base.Model
	Title   string //资讯标题
	Des     string //描述
	Content string //资讯内容
}
