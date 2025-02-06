package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(host, port, user, pass, dbname string) {
	ConnectToMySQL(host, port, user, pass, dbname)
	DB.Logger = logger.Default.LogMode(logger.Info)
}

func ConnectToMySQL(host, port, user, pass, dbname string) error {
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = d
	return nil
}

func CloseMySQL() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
