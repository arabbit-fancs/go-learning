package service

import (
	"ctag-go-learning/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db = newDBConn()

func newDBConn() *gorm.DB {
	db, err := gorm.Open("mysql", conf.Env("dbUser")+":"+conf.Env("dbPass")+"@tcp("+conf.Env("dbHost")+")/"+conf.Env("dbName")+"?charset=utf8mb4&parseTime=true&collation=utf8_general_ci&interpolateParams=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}
	return db
}

func GetDBCon() *gorm.DB {
	return db
}

func CloseDBCon() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
