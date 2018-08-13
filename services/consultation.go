package services

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/models"
	"github.com/MEIGUOSHU/meiguoshu_api_server/common"
)

func CreateConsultation(consu *models.Consultation) (*models.Consultation, error) {
	db := common.DB
	res, err := consu.Create(db)
	return res, err
}

func UpdateConsultationById(consu *models.Consultation) (*models.Consultation, error) {
	res, err := consu.Update(common.DB)
	return res, err
}

func DeleteConsultationById(consu *models.Consultation) error {
	db := common.DB
	err := consu.Delete(db)
	return err
}

func GetOneConsultation(consu *models.Consultation) (*models.Consultation, error) {
	var model models.Consultation
	db := common.DB
	err := db.Where("id = ?", consu.ID).First(&model).Error
	return &model, err
}

func GetAllConsultation(whereQuery string, whereArgs []interface{}, page int, limit int, order string) ([]*models.Consultation, int, error) {
	var (
		err	error
		total int
		list []*models.Consultation
	)
	db := common.DB

	err = db.Where(whereQuery, whereArgs...).Offset((page - 1) * limit).Order(order).Limit(limit).Find(&list).Count(&total).Error
	return list, total, err
}
