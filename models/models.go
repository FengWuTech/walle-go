package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"walle-go/conf"
	"walle-go/logger"
)

const (
	StatusRemove = -1
	StatusDefault = 1
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

// Setup initializes the database instance
func Setup(debug bool) {
	var (
		err error
		c conf.Conf
	)
	c.GetConf()
	db, err = gorm.Open("mysql", c.Db)

	if err != nil {
		logger.Fatalf("models.Setup err: %v", err.Error())
	}

	db.LogMode(debug)

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Set("gorm:association", false)                //禁止自动创建/更新包含关系
	db.Set("gorm:association_save_reference", false) //禁止自动创建关联关系
}
