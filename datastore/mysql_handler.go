package datastore

import (
	"WALLET-AS-A-SERVICE/config"
	"WALLET-AS-A-SERVICE/dao/entity"
	"WALLET-AS-A-SERVICE/utils"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB
var err error

func InitializeDb() {
	Dsn := config.GetAllConfiguration()
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               Dsn, // data source name
		DefaultStringSize: 256}), &gorm.Config{})
	utils.CheckError(err)
	DB.AutoMigrate(&entity.User{}, &entity.Wallet{}, &entity.Transaction{})
}

func GetDbHandle() *gorm.DB {
	return DB
}
