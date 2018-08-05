package model

import (
	"github.com/MEIGUOSHU/web-api/packages/model/base"
)

type Link struct {
	base.Model
	Img  string
	Name string
	Url  string
}
