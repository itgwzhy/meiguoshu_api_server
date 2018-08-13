package models

import "github.com/jinzhu/gorm"

// 品牌资讯
type News struct {
	BaseModel
	Title   string //资讯标题
	Des     string //描述
	Content string //资讯内容
}

func (new *News) Create(db *gorm.DB) (*News, error) {
	var model *News
	err := db.Create(&model).Error
	db.First(&model)
	return model, err
}

func (new *News) Update(db *gorm.DB) (*News, error) {
	var model *News
	err := db.Updates(&model).Error
	db.First(&model)
	return model, err
}

func (new *News) Delete(db *gorm.DB) (error) {
	err := db.Delete(&new).Error
	return err
}
