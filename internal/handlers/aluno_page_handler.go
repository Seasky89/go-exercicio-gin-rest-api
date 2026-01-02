package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlunoPageHandler interface {
	IndexPage(c *gin.Context)
	NotFoundPage(c *gin.Context)
}

func (h *alunoHandler) IndexPage(c *gin.Context) {
	alunos, err := h.service.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", gin.H{
			"error": "Erro ao carregar os alunos",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func (h *alunoHandler) NotFoundPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
