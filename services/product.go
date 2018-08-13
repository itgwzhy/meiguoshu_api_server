package services

import "github.com/MEIGUOSHU/meiguoshu_api_server/models"
import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/common"
)

func CreateProduct(pro *models.Product) (*models.Product, error) {
	db := common.DB
	res, err := pro.Create(db)
	return res, err
}


func UpdateProduct(pro *models.Product) (*models.Product, error) {
	res, err := pro.Update(common.DB)
	return res, err
}

func DeleteProductById(pro *models.Product) error {
	db := common.DB
	err := pro.Delete(db)
	return err
}

func GetProductById(news *models.Product) (*models.Product, error) {
	var model models.Product
	db := common.DB
	err := db.Where("id = ?", news.ID).First(&model).Error
	return &model, err
}

func GetAllProduct(whereQuery string, whereArgs []interface{}, page int, limit int, order string) ([]*models.Product, int, error) {
	var (
		err	error
		total int
		list []*models.Product
	)
	db := common.DB

	err = db.Where(whereQuery, whereArgs...).Offset((page - 1) * limit).Order(order).Limit(limit).Find(&list).Count(&total).Error
	return list, total, err
}
