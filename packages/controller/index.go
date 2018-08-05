package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/MEIGUOSHU/web-api/utils"
	"github.com/MEIGUOSHU/web-api/tools/msg"
)

type Index struct{}

func (index *Index) Get(c *gin.Context) {
	var code = 200

	id := c.Param("id")

	if id == "1" {

	}


	//返回信息
	utils.Res(c, code, nil, msg.GetResponseMessage(code))
}

func (index *Index) Post() {

}

func (index *Index) Put() {

}

func (index *Index) Delete() {

}

func (index *Index) GetList() {

}
