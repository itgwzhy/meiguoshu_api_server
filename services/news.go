package services

import "github.com/MEIGUOSHU/meiguoshu_api_server/models"
import "github.com/MEIGUOSHU/meiguoshu_api_server/common"

func CreateNews(news *models.News) (*models.News, error) {
	db := common.DB
	res, err := news.Create(db)
	return res, err
}


func UpdateNews(news *models.News) (*models.News, error) {
	res, err := news.Update(common.DB)
	return res, err
}

func DeleteNewsById(news *models.News) error {
	db := common.DB
	err := news.Delete(db)
	return err
}

func GetNewsById(news *models.News) (*models.News, error) {
	var model models.News
	db := common.DB
	err := db.Where("id = ?", news.ID).First(&model).Error
	return &model, err
}

func GetAllNews(whereQuery string, whereArgs []interface{}, page int, limit int, order string) ([]*models.News, int, error) {
	var (
		err	error
		total int
		list []*models.News
	)
	db := common.DB

	err = db.Where(whereQuery, whereArgs...).Offset((page - 1) * limit).Order(order).Limit(limit).Find(&list).Count(&total).Error
	return list, total, err
}
