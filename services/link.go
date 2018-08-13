package services

import "github.com/MEIGUOSHU/meiguoshu_api_server/models"
import "github.com/MEIGUOSHU/meiguoshu_api_server/common"

func CreateLink(link *models.Link) (*models.Link, error) {
	db := common.DB
	res, err := link.Create(db)
	return res, err
}


func UpdateLink(link *models.Link) (*models.Link, error) {
	res, err := link.Update(common.DB)
	return res, err
}

func DeleteLinkById(link *models.Link) error {
	db := common.DB
	err := link.Delete(db)
	return err
}

func GetLinkById(link *models.Link) (*models.Link, error) {
	var model models.Link
	db := common.DB
	err := db.Where("id = ?", link.ID).First(&model).Error
	return &model, err
}

func GetAllLink(whereQuery string, whereArgs []interface{}, page int, limit int, order string) ([]*models.Link, int, error) {
	var (
		err	error
		total int
		list []*models.Link
	)
	db := common.DB

	err = db.Where(whereQuery, whereArgs...).Offset((page - 1) * limit).Order(order).Limit(limit).Find(&list).Count(&total).Error
	return list, total, err
}
