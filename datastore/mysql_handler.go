package datastore

import (
	"WALLET-AS-A-SERVICE/config"
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/utils"
	"WALLET-AS-A-SERVICE/vendor/gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitializeDb() {
	DSN := config.GetAllConfiuration()
	DB, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	utils.CheckError(err)
	DB.AutoMigrate(&entity.User{}, &entity.Wallet{}, &entity.Transaction{})
}

func GetDbHandle() *gorm.DB {
	return DB
}
