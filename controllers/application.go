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

/**
 * @api {get} /user/:id Request User information
 * @apiName GetUser
 * @apiGroup User
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname  Lastname of the User.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "firstname": "John",
 *       "lastname": "Doe"
 *     }
 *
 * @apiError UserNotFound The id of the User was not found.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error": "UserNotFound"
 *     }
 */
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
func UpdateApplication(ctx iris.Context) {
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
func GetOneApplication(ctx iris.Context) {
	var (
		app models.Application
		reApp *models.Application
		err	error
	)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数错误")
		return
	}
	app.ID = id
	reApp, err = services.GetApplicationById(&app)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "查询失败")
	}else{
		common.Responese(ctx, http.StatusOK, reApp, "查询成功")
	}
}

//综合查询
func GetAllApplication(ctx iris.Context) {
	var (
		page int
		limit int
		order string
		whereQuery  string
		whereArgs   []interface{}
		list []*models.Application
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

	list, total, err = services.GetAllApplication(whereQuery, whereArgs, page, limit, order)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "查询列表失败")
	}else{
		data["data_list"] = list
		data["count"] = total
		common.Responese(ctx, http.StatusOK, data, "查询成功")
	}
}
