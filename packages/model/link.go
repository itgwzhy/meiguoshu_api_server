package model

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/packages/model/base"
)

type Link struct {
	base.Model
	Img  string
	Name string
	Url  string
}
