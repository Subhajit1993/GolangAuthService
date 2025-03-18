package modules

import (
	"github.com/gin-gonic/gin"
	"os"
)

func HealthCheck(context *gin.Context) {
	file, _ := os.Open("nonexistentfile.txt")
	defer file.Close()
	context.JSON(200, gin.H{
		"message": "OK",
	})
}
