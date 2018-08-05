package controller

import (
	"github.com/MEIGUOSHU/web-api/tools/msg"
	"github.com/MEIGUOSHU/web-api/utils"
	"github.com/gin-gonic/gin"
)

type Product struct{}

func (pro *Product) Get(c *gin.Context) {
	var code = 200
	id := c.Param("id")
	if id == "1" {}
	//返回信息
	utils.Res(c, code, nil, msg.GetResponseMessage(code))
}

func (pro *Product) Post(c *gin.Context) {

}

func (pro *Product) Put(c *gin.Context) {

}

func (pro *Product) Delete(c *gin.Context) {

}

func (pro *Product) GetList(c *gin.Context) {

}
