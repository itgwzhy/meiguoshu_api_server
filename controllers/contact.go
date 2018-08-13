package controllers

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/common"
	"github.com/MEIGUOSHU/meiguoshu_api_server/models"
	"github.com/MEIGUOSHU/meiguoshu_api_server/services"
	"github.com/kataras/iris"
	"net/http"
	"strconv"
	"fmt"
)

func CreateContact(ctx iris.Context) {
	var app models.Contact

	if err := ctx.ReadJSON(&app); err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数错误")
		return
	}

	res, err := services.CreateContact(&app)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "创建数据失败")
	} else {
		common.Responese(ctx, http.StatusOK, res, "创建数据成功！")
	}
}

//编辑
func UpdateContact(ctx iris.Context) {
	var app models.Contact
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
	res, err := services.UpdateContact(&app)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "修改数据失败")
	} else {
		common.Responese(ctx, http.StatusOK, res, "修改数据成功")
	}
}

//删除
func DeleteContactById(ctx iris.Context) {
	var app models.Contact
	id := ctx.Params().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数类型转换错误")
		return
	}
	app.ID = idInt
	err = services.DeleteContactById(&app)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "删除数据失败")
	} else {
		common.Responese(ctx, http.StatusOK, nil, "删除数据成功")
	}
}

//按ID查询
func GetOneContact(ctx iris.Context) {
	var (
		app models.Contact
		reApp *models.Contact
		err	error
	)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数错误")
		return
	}
	app.ID = id
	reApp, err = services.GetContactById(&app)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "查询失败")
	}else{
		common.Responese(ctx, http.StatusOK, reApp, "查询成功")
	}
}

//综合查询
func GetAllContact(ctx iris.Context) {
	var (
		page int
		limit int
		order string
		whereQuery  string
		whereArgs   []interface{}
		list []*models.Contact
		total int
		err error
	)
	data := make(map[string]interface{})

	page = ctx.URLParamIntDefault("page", 1)
	limit = ctx.URLParamIntDefault("limit", 15)
	order = ctx.URLParamDefault("order", "id desc")

	whereQuery += "1 = ?"
	whereArgs = append(whereArgs, 1)

	if name := ctx.URLParam("name"); name != "" {
		whereQuery += " AND name ILIKE ? "
		whereArgs = append(whereArgs, fmt.Sprintf("%%%s%%", name))
	}

	list, total, err = services.GetAllContact(whereQuery, whereArgs, page, limit, order)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "查询列表失败")
	}else{
		data["data_list"] = list
		data["count"] = total
		common.Responese(ctx, http.StatusOK, data, "查询成功")
	}
}
