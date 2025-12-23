package domain

import "errors"

var (
	ErrNomeObrigatorio = errors.New("nome é obrigatório")
	ErrCpfObrigatorio  = errors.New("CPF é obrigatório")
	ErrCpfTamanho      = errors.New("CPF deve contar 11 números")
	ErrCpfInvalido     = errors.New("CPF inválido")
	ErrRgObrigatorio   = errors.New("RG é obrigarório")
	ErrNotFound        = errors.New("aluno não encontrado")
)
