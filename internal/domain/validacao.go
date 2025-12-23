package domain

import (
	"strconv"

	"github.com/Seasky89/go-gin-rest-api/internal/models"
)

func ValidarAluno(a *models.Aluno) error {
	if a.Nome == "" {
		return ErrNomeObrigatorio
	}

	if a.CPF == "" {
		return ErrCpfObrigatorio
	}

	if len(a.CPF) != 11 {
		return ErrCpfTamanho
	}

	if _, err := strconv.Atoi(a.CPF); err != nil {
		return ErrCpfInvalido
	}

	return nil
}
