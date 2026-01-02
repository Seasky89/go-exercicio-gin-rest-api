package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(CustomRecovery())

	if gin.Mode() != gin.TestMode {
		r.LoadHTMLGlob("templates/*")
		r.Static("/assets", "./assets")
	}

	return r
}
