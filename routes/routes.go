package routes

import (
	"github.com/Seasky89/go-gin-rest-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func HandleRequests(h *handlers.AlunoHandler) {
	r := gin.Default()
	r.GET("/:nome", h.Welcome)
	r.GET("/alunos", h.FindAll)
	r.GET("/alunos/:id", h.FindById)
	r.POST("/alunos", h.Create)
	r.DELETE("/alunos/:id", h.Delete)
	r.PATCH("/alunos/:id", h.Update)
	r.GET("/alunos/cpf/:cpf", h.FindByCpf)
	r.Run()
}
