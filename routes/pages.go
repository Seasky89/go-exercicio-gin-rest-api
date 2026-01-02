package routes

import (
	"github.com/Seasky89/go-gin-rest-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterPageRoutes(r *gin.Engine, h handlers.AlunoPageHandler) {
	r.GET("/index", h.IndexPage)
	r.NoRoute(h.NotFoundPage)
}
