package router

import (
	"ctag-go-learning/cont"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(cors)
	r.GET("/alive", func(c *gin.Context) {
		c.Status(200)
		return
	})

	r.POST("/product", cont.Create)
	r.GET("/product/:id", cont.Get)
	r.POST("/product/:id", cont.Update)
	r.DELETE("/product/:id", cont.Delete)

	return r
}

func cors(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", headers)

	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}
	c.Next()
}
