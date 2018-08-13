package models

import "github.com/jinzhu/gorm"

type Link struct {
	BaseModel
	Img  string
	Name string
	Url  string
}

func (link *Link) Create(db *gorm.DB) (*Link, error) {
	var model *Link
	err := db.Create(&model).Error
	db.First(&model)
	return model, err
}

func (link *Link) Update(db *gorm.DB) (*Link, error) {
	var model *Link
	err := db.Updates(&model).Error
	db.First(&model)
	return model, err
}

func (link *Link) Delete(db *gorm.DB) (error) {
	err := db.Delete(&link).Error
	return err
}
