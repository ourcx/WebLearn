package config

import (
	"example/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func initDB() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//你链接的数据库是什么,后面他就去找对应的数据库,不会去找你的其他表
	//ab是对应的数据库链接，err是可能发生的错误
	if err != nil {
		log.Fatal("链接数据库失败，请重新链接：%v", err)
	} else {
		fmt.Print("链接成功")
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatal(err)
	}

	global.Db = db
}

//db *gorm.DB
