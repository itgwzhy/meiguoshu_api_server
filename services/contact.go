package services

import "github.com/MEIGUOSHU/meiguoshu_api_server/models"
import "github.com/MEIGUOSHU/meiguoshu_api_server/common"

func CreateContact(contact *models.Contact) (*models.Contact, error) {
	db := common.DB
	res, err := contact.Create(db)
	return res, err
}

func UpdateContact(contact *models.Contact) (*models.Contact, error) {
	res, err := contact.Update(common.DB)
	return res, err
}

func DeleteContactById(contact *models.Contact) error {
	db := common.DB
	err := contact.Delete(db)
	return err
}

func GetContactById(contact *models.Contact) (*models.Contact, error) {
	var model models.Contact
	db := common.DB
	err := db.Where("id = ?", contact.ID).First(&model).Error
	return &model, err
}

func GetAllContact(whereQuery string, whereArgs []interface{}, page int, limit int, order string) ([]*models.Contact, int, error) {
	var (
		err	error
		total int
		list []*models.Contact
	)
	db := common.DB

	err = db.Where(whereQuery, whereArgs...).Offset((page - 1) * limit).Order(order).Limit(limit).Find(&list).Count(&total).Error
	return list, total, err
}