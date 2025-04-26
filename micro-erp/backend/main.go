package main

import (
	"micro-erp/config"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	sqlDB, _ := db.db
	defer sqlDB.Close()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Methods == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.setupRouts(r)

	r.Run(":8080")
}
