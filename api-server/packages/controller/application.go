package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/MEIGUOSHU/api-server/utils"
)

type Application struct{}

func (app *Application) Get(c *gin.Context) {
	//id := utils.Str2Int(c.Param("id"))
	code := 200
	utils.Res(c, code, nil, "success")
}

func (app *Application) POST(c *gin.Context) {

}

func (app *Application) PUT(c *gin.Context) {

}

func (app *Application) DELETE(c *gin.Context) {

}

func (app *Application) GetList(c *gin.Context) {

}

