package system

import (
	"github.com/GACHAIN/dongjing-server/models"
	"github.com/jinzhu/gorm"
)

func DBMigrate(db *gorm.DB) {
	//migrate base model to database table
	db.AutoMigrate(&models.User{}, &models.Menu{})
	//application model to database table
	db.AutoMigrate()
}
