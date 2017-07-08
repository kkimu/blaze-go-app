package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Db Db
var db *gorm.DB

func init() {
	ip := os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	var err error
	var host string
	if ip != "" {
		host = "root:@tcp(" + ip + ":3306)/blaze?parseTime=True&loc=Japan"
	} else {
		host = "root:@tcp(127.0.0.1:3306)/blaze?parseTime=True&loc=Japan"
	}
	db, err = gorm.Open("mysql", host)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.DB()
	db.AutoMigrate(&Video{})
}
