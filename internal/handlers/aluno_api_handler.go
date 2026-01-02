package handlers

import (
	"net/http"
	"strconv"

	"github.com/Seasky89/go-gin-rest-api/internal/dto"
	"github.com/Seasky89/go-gin-rest-api/internal/models"
	"github.com/Seasky89/go-gin-rest-api/internal/services"
	"github.com/gin-gonic/gin"
)

type AlunoAPIHandler interface {
	Welcome(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	FindByCpf(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
}

type alunoHandler struct {
	service services.AlunoService
}

func NewAlunoHandler(service services.AlunoService) *alunoHandler {
	return &alunoHandler{service: service}
}

func (h *alunoHandler) Welcome(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func (h *alunoHandler) FindAll(c *gin.Context) {
	a, err := h.service.FindAll()
	if err != nil {
		HandleHttpError(c, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *alunoHandler) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	a, err := h.service.FindById(id)
	if err != nil {
		HandleHttpError(c, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *alunoHandler) FindByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	a, err := h.service.FindByCpf(cpf)
	if err != nil {
		HandleHttpError(c, err)
		return
	}
	c.JSON(http.StatusOK, a)

}

func (h *alunoHandler) Create(c *gin.Context) {
	var req dto.CreateAlunoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleHttpError(c, err)
		return
	}

	aluno := models.Aluno{
		Nome: req.Nome,
		CPF:  req.CPF,
		RG:   req.RG,
	}

	created, err := h.service.Create(aluno)
	if err != nil {
		HandleHttpError(c, err)
		return
	}

	resp := dto.AlunoResponse{
		ID:   created.ID,
		Nome: created.Nome,
		CPF:  created.CPF,
		RG:   created.RG,
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *alunoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}
	a, err := h.service.Delete(id)
	if err != nil {
		HandleHttpError(c, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *alunoHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	var req dto.UpdateAlunoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleHttpError(c, err)
		return
	}

	updated, err := h.service.Update(id, req)
	if err != nil {
		HandleHttpError(c, err)
		return
	}

	resp := dto.AlunoResponse{
		ID:   updated.ID,
		Nome: updated.Nome,
		CPF:  updated.CPF,
		RG:   updated.RG,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *alunoHandler) Patch(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id inválido",
		})
		return
	}

	var req dto.PatchAlunoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleHttpError(c, err)
		return
	}

	patched, err := h.service.Patch(id, req)
	if err != nil {
		HandleHttpError(c, err)
		return
	}

	resp := dto.AlunoResponse{
		ID:   patched.ID,
		Nome: patched.Nome,
		CPF:  patched.CPF,
		RG:   patched.RG,
	}

	c.JSON(http.StatusOK, resp)
}
