package system

import (
	"github.com/jinzhu/gorm"
	"github.com/MEIGUOSHU/meiguoshu_api_server/packages/model"
)

func DBMigrate(db *gorm.DB) {
	//application model to database table
	db.AutoMigrate(
		&model.Application{},
		&model.Category{},
		&model.Consultation{},
		&model.Contact{},
		&model.Link{},
		&model.News{},
		&model.Product{})
}
