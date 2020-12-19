package mappings

import (
	// "github.com/gin-goinc/gin"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func CreateUrlMappings() {
	R = gin.Default()
	api := R.Group("/api")
	{

		api.GET("/test", func(resp *gin.Context){
			resp.JSON(200, gin.H{
				"message": "test successful",
			});
		})
	}
}