package handlers

import (
	"net/http"
	"strconv"

	"github.com/Seasky89/go-gin-rest-api/internal/dto"
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
	a, err := h.service.FindAll()
	if err != nil {
		HandleHttpError(c, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *AlunoHandler) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		HandleHttpError(c, err)
		return
	}

	a, err := h.service.FindById(id)
	if err != nil {
		HandleHttpError(c, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *AlunoHandler) FindByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	a, err := h.service.FindByCpf(cpf)
	if err != nil {
		HandleHttpError(c, err)
		return
	}
	c.JSON(http.StatusOK, a)

}

func (h *AlunoHandler) Create(c *gin.Context) {
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

func (h *AlunoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		HandleHttpError(c, err)
		return
	}
	a, err := h.service.Delete(id)
	if err != nil {
		HandleHttpError(c, err)
		return
	}
	c.JSON(http.StatusOK, a)
}

func (h *AlunoHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		HandleHttpError(c, err)
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

func (h *AlunoHandler) Patch(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		HandleHttpError(c, err)
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
