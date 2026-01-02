package routes

import (
	"github.com/Seasky89/go-gin-rest-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine, h handlers.AlunoAPIHandler) {
	r.GET("/welcome/:nome", h.Welcome)
	alunos := r.Group("/alunos")
	{
		alunos.GET("", h.FindAll)
		alunos.GET("/:id", h.FindById)
		alunos.POST("", h.Create)
		alunos.DELETE("/:id", h.Delete)
		alunos.PUT("/:id", h.Update)
		alunos.PATCH("/:id", h.Patch)
		alunos.GET("/cpf/:cpf", h.FindByCpf)
	}
}
