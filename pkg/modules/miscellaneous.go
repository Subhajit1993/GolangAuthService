package modules

import "github.com/gin-gonic/gin"

func SampleFunction(a string, b string, c string) string {
	return a + b
}

func HealthCheck(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "OK",
	})
}
