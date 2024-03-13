package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"palworld/rpc/goods/model"
)

var Db *gorm.DB

func InitDB(dataSource string) error {
	if Db != nil {
		return nil
	}
	db, err := NewDB(dataSource)
	if err != nil {
		return err
	}
	Db = db
	SetDefault(db)
	return nil
}

func NewDB(dataSource string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Goods{})
	log.Println("数据库连接成功")
	return db, nil
}
