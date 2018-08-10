package services

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/common"
	"github.com/MEIGUOSHU/meiguoshu_api_server/models"
)

func CreateApplication(app *models.Application) (*models.Application, error) {
	res, err := app.Create(common.DB)
	return res, err
}

func UpdateApplication(app *models.Application) (*models.Application, error) {
	res, err := app.Update(common.DB)
	return res, err
}

func DeleteApplicationById(app *models.Application) (*models.Application, error) {
	nil, err := app.Delete(common.DB)
	return nil, err
}

func GetApplicationById(app *models.Application) (*models.Application, error) {
	var model models.Application
	db := common.DB
	id := app.ID
	err := db.Where("id = ?", id).First(&model).Error
	return &model, err
}
