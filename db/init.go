package db

import (
	"IceBreaking/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Get() *gorm.DB {
	return db
}

func init() {
	DbConn()
	creatTable()
}

// 创建数据库连接
func DbConn()  {
	c := config.Get()
	m := c.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s" +
		"?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Pwd, m.Host, m.Port, m.Database)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

//表不存在则创建
func creatTable()  {
	db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(&Student{}, &Picture{}, &AssStuPic{})
}
