package models

import "github.com/jinzhu/gorm"

//联系方式
type Contact struct {
	BaseModel
	Tel     string //联系电话
	Fax     string //传真
	Working string
	Email   string //邮箱
	Address string //地址
	Weibo   string //微博
	Qrcode  string //二维码
}

func (cotact *Contact) Create(db *gorm.DB) (*Contact, error) {
	var model *Contact
	err := db.Create(&model).Error
	db.First(&model)
	return model, err
}

func (cotact *Contact) Update(db *gorm.DB) (*Contact, error) {
	var model *Contact
	err := db.Updates(&model).Error
	db.First(&model)
	return model, err
}

func (cotact *Contact) Delete(db *gorm.DB) error {
	err := db.Delete(&cotact).Error
	return err
}
