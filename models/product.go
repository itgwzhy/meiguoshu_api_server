package models

import "github.com/jinzhu/gorm"

//产品中心
type Product struct {
	BaseModel
	Name        string `json:"name"`     //产品名称
	Number      int64  `json:"number"`   //产品编号
	Spec        string `json:"spec"`     //商品规格
	Qgp         string `json:"qgp"`      //保质期
	ViewNum     int64  `json:"view_num"` //浏览次数
	Content     string `json:"content"`  //商品详情
	Category_id int64  `json:"category_id"`
}

func (pro *Product) Create(db *gorm.DB) (*Product, error) {
	var model *Product
	err := db.Create(&pro).Error
	if err == nil {
		db.Where("id = ?", pro.ID).First(&model)
	}

	return pro, err
}

func (pro *Product) Update(db *gorm.DB) (*Product, error) {
	var model *Product
	err := db.Updates(&model).Error
	db.First(&model)
	return model, err
}

func (pro *Product) Delete(db *gorm.DB) (error) {
	err := db.Delete(&pro).Error
	return err
}
