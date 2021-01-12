package cont

import (
	"ctag-go-learning/service"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	service.Production.ExampleFunction(1)
}

func Get(c *gin.Context) {
}

func Update(c *gin.Context) {
}

func Delete(c *gin.Context) {
}