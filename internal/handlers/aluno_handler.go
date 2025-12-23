package handlers

import (
	"net/http"
	"strconv"

	"github.com/Seasky89/go-gin-rest-api/internal/models"
	"github.com/Seasky89/go-gin-rest-api/internal/services"
	"github.com/gin-gonic/gin"
)

type AlunoHandler struct {
	service *services.AlunoService
}

func NewAlunoHandler(service *services.AlunoService) *AlunoHandler {
	return &AlunoHandler{service: service}
}

func (h *AlunoHandler) Welcome(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func (h *AlunoHandler) FindAll(c *gin.Context) {
	var a []models.Aluno

	a, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *AlunoHandler) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	a, err := h.service.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *AlunoHandler) FindByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	a, err := h.service.FindByCpf(cpf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, a)

}

func (h *AlunoHandler) Create(c *gin.Context) {
	var data models.Aluno
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	a, err := h.service.Create(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

func (h *AlunoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	a, err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *AlunoHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var data models.Aluno
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	a, err := h.service.Update(id, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, a)
}
