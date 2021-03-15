package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//初始化数据库
//var DB *sql.DB
var DB *gorm.DB



func InitDB(){
	dsn := "host=localhost user=postgres password=123 dbname=go_mars port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		panic("数据库连接失败")
	}
	sqlDB,err :=DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//自动生成表
	DB.AutoMigrate(&User{},&Article{},&Tag{},&ArticleTag{})
}
