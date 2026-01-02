package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Println("PANIC:", recovered)
		c.HTML(500, "500.html", nil)
	})
}
