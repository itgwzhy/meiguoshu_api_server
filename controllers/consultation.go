package controllers

import (
	"github.com/kataras/iris"
	"github.com/MEIGUOSHU/meiguoshu_api_server/models"
	"github.com/MEIGUOSHU/meiguoshu_api_server/common"
	"net/http"
	"github.com/MEIGUOSHU/meiguoshu_api_server/services"
	"fmt"
)

//创建
func CreateConsultation(ctx iris.Context) {
	var (
		res *models.Consultation
		err error
	)

	var consu models.Consultation
	err = ctx.ReadJSON(&consu)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "参数绑定失败")
		return
	}
	res, err = services.CreateConsultation(&consu)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "创建数据失败")
	}else{
		common.Responese(ctx, http.StatusOK, res, "创建数据成功")
	}
}

//编辑
func UpdateConsultation(ctx iris.Context) {
	//1.接收ID 和 修改的参数
	var (
		consu models.Consultation
		res *models.Consultation
		err error
	)

	id, err := ctx.Params().GetInt("id")
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "ID获取失败")
		return
	}
	err = ctx.ReadJSON(&consu)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "接收修改参数失败")
		return
	}
	consu.ID = id
	//2.修改数据
	res, err = services.UpdateConsultationById(&consu)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "更新数据失败")
	}else{
		common.Responese(ctx, http.StatusOK, res, "更新数据成功")
	}
}


//删除
func DeleteConsultationById(ctx iris.Context) {
	var (
		consu models.Consultation
	)
	//1.接收ID
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "ID获取失败")
		return
	}
	consu.ID = id
	//2.删除
	err = services.DeleteConsultationById(&consu)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "删除失败")
	}else{
		common.Responese(ctx, http.StatusOK, err.Error(), "删除成功")
	}
}

//GetOne
func GetOneConsultation(ctx iris.Context) {
	var (
		consu models.Consultation
	)
	//1.接收ID
	id, err := ctx.Params().GetInt("id")
	//2.查询
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "ID获取失败")
		return
	}
	consu.ID = id
	res, err := services.GetOneConsultation(&consu)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "查询失败")
	}else{
		common.Responese(ctx, http.StatusOK, res, "查询成功")
	}
}

//GetAll
func GetAllConsultation(ctx iris.Context) {
	var (
		page int
		limit int
		order string
		whereQuery  string
		whereArgs   []interface{}
		list []*models.Consultation
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

	list, total, err = services.GetAllConsultation(whereQuery, whereArgs, page, limit, order)
	if err != nil {
		common.Responese(ctx, http.StatusForbidden, err.Error(), "查询列表失败")
	}else{
		data["data_list"] = list
		data["count"] = total
		common.Responese(ctx, http.StatusOK, data, "查询成功")
	}
}