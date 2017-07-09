package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Db Db
var db *gorm.DB

func init() {
	var err error
	var host string
	host = "root:@tcp(spajamserver_mysql_1:3306)/blaze?parseTime=True&loc=Japan"
	db, err = gorm.Open("mysql", host)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&Video{})
	db.DB()
	//db.AutoMigrate(&Video{})
}
