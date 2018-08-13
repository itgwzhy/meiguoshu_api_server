package models

import "github.com/jinzhu/gorm"

//分类
type Category struct {
	BaseModel
	name  string //名称
	tty   int64  //类型 1：产品中心分类 2：研发推广分类
	pid   int64  //父iD
	level int64  //公司等级
}

func (cat *Category) Create(db *gorm.DB) (*Category, error) {
	var model *Category
	err := db.Create(&model).Error
	db.First(&model)
	return model, err
}

func (cat *Category) Update(db *gorm.DB) (*Category, error) {
	var model *Category
	err := db.Updates(&model).Error
	db.First(&model)
	return model, err
}

func (cat *Category) Delete(db *gorm.DB) (*Category, error) {
	err := db.Delete(&cat).Error
	return nil, err
}
