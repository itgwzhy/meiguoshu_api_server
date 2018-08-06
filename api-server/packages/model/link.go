package model

import (
	"github.com/MEIGUOSHU/api-server/packages/model/base"
)

type Link struct {
	base.Model
	Img  string
	Name string
	Url  string
}
