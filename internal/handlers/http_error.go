package handlers

import (
	"errors"
	"net/http"

	"github.com/Seasky89/go-gin-rest-api/internal/domain"
	"github.com/gin-gonic/gin"
)

func HandleHttpError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

	case errors.Is(err, domain.ErrNomeObrigatorio),
		errors.Is(err, domain.ErrCpfObrigatorio),
		errors.Is(err, domain.ErrCpfTamanho),
		errors.Is(err, domain.ErrCpfInvalido),
		errors.Is(err, domain.ErrRgObrigatorio):
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno"})
	}
}
