package main

import (
	"ctag-go-learning/model"
	"ctag-go-learning/service"
)

func main() {
	db := service.GetDBCon()

	db.DropTableIfExists(&model.Example{})
	db.AutoMigrate(&model.Example{})


	defer service.CloseDBCon()
}
