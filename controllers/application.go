package controllers

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/common"
	"github.com/MEIGUOSHU/meiguoshu_api_server/models"
	"github.com/MEIGUOSHU/meiguoshu_api_server/services"
	"github.com/kataras/iris"
	"net/http"
	"strconv"
)

//创建
func CreateApplication(ctx iris.Context) {
	var app models.Application

	if err := ctx.ReadJSON(&app); err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数错误")
		return
	}

	res, err := services.CreateApplication(&app)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "创建数据失败")
	} else {
		common.Responese(ctx, http.StatusOK, res, "创建数据成功！")
	}
}

//编辑
func UpdateApplicationById(ctx iris.Context) {
	var app models.Application
	//接收ID
	id := ctx.Params().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数类型转换错误")
		return
	}
	if err := ctx.ReadJSON(&app); err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数错误")
		return
	}
	app.ID = idInt
	res, err := services.UpdateApplication(&app)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "修改数据失败")
	} else {
		common.Responese(ctx, http.StatusOK, res, "修改数据成功")
	}
}

//删除
func DeleteApplicationById(ctx iris.Context) {
	var app models.Application
	id := ctx.Params().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数类型转换错误")
		return
	}
	app.ID = idInt
	_, err = services.DeleteApplicationById(&app)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "删除数据失败")
	} else {
		common.Responese(ctx, http.StatusOK, nil, "删除数据成功")
	}
}

//按ID查询
func GetApplicationById(ctx iris.Context) {
	var app models.Application
	id := ctx.Params().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数错误")
		return
	}
	app.ID = idInt
	services.GetApplicationById(&app)
}

//综合查询
func GetApplicationAll(ctx iris.Context) {
	ctx.Writef("this is get apps")
}
