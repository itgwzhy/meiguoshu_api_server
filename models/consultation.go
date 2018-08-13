package models

import "github.com/jinzhu/gorm"

//咨询
type Consultation struct {
	BaseModel
	Name    string //称呼
	City    string //城市
	Company string //公司
	Mobile  string //手机号码
	Email   string //邮箱
	Content string //咨询内容
}

func (cstion *Consultation) Create(db *gorm.DB) (*Consultation, error) {
	var model *Consultation
	err := db.Create(&model).Error
	db.First(&model)
	return model, err
}

func (cstion *Consultation) Update(db *gorm.DB) (*Consultation, error) {
	var model *Consultation
	err := db.Updates(&model).Error
	db.First(&model)
	return model, err
}

func (cstion *Consultation) Delete(db *gorm.DB) error {
	err := db.Delete(&cstion).Error
	return err
}
