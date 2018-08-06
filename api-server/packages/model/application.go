package model

import (
	"github.com/MEIGUOSHU/api-server/packages/model/base"
)

//研发推广
type Application struct {
	base.Model
	Name  string 	//名称
	Code	string //编码
	Image string	//图片地址
	Use   string	//用了啥
	Talk  string	//大师说
	Eat   string	//何时吃
}
