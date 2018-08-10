package models

import (
	"github.com/jinzhu/gorm"
)

//研发推广
type Application struct {
	BaseModel
	Name  string `json:"name"`  //名称
	Code  int    `json:"code"`  //编码
	Image string `json:"image"` //图片地址
	Use   string `json:"use"`   //用了啥
	Talk  string `json:"talk"`  //大师说
	Eat   string `json:"eat"`   //何时吃
}

func (app *Application) Create(db *gorm.DB) (*Application, error) {
	var model *Application
	err := db.Create(&model).Error
	db.First(&model)
	return model, err
}

func (app *Application) Update(db *gorm.DB) (*Application, error) {
	var model *Application
	err := db.Updates(&model).Error
	db.First(&model)
	return model, err
}

func (app *Application) Delete(db *gorm.DB) (*Application, error) {
	err := db.Delete(&app).Error
	return nil, err
}
