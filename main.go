package main

import (
	// "github.com/gin-gonic/gin"
	"gin-framework/mappings"
	// "github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin"
)

// var R *gin.Engine

func main(){
	mappings.CreateUrlMappings()

	// R = gin.Default()
	// api := R.Group("/api")
	// {
	// 	api.GET("/test", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"message": "test successful",
	// 		}) 
	// 	})
	// }

	mappings.R.Run(":5000")
}