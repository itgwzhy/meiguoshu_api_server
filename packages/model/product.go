package model

import (
	"github.com/MEIGUOSHU/web-api/packages/model/base"
)

//产品中心
type Product struct {
	base.Model
	Name        string `json:"name"`     //产品名称
	Number      int64  `json:"number"`   //产品编号
	Spec        string `json:"spec"`     //商品规格
	Qgp         string `json:"qgp"`      //保质期
	ViewNum     int64  `json:"view_num"` //浏览次数
	Content     string `json:"content"`  //商品详情
	Category_id int64  `json:"category_id"`
}
