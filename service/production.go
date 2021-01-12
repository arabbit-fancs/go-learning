package service

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

var Production = newProductionService()

type productionService struct {
	db *gorm.DB
}

func newProductionService() productionService {
	return productionService{db: db}
}

func (p *productionService) ExampleFunction(id int) string {

	return strconv.Itoa(id)
}
