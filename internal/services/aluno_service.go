package services

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Seasky89/go-gin-rest-api/internal/models"
	"github.com/Seasky89/go-gin-rest-api/internal/repository"
)

type AlunoService struct {
	repo *repository.AlunoRepository
}

var (
	ErrNomeObrigatorio = errors.New("nome é obrigatório")
	ErrCpfObrigatorio  = errors.New("CPF é obrigatório")
	ErrCpfTamanho      = errors.New("CPF deve contar 11 números")
	ErrCpfInvalido     = errors.New("CPF inválido")
	ErrRgObrigatorio   = errors.New("RG é obrigarório")
	ErrNotFound        = errors.New("aluno não encontrado")
)

func NewAlunoService(repo *repository.AlunoRepository) *AlunoService {
	return &AlunoService{repo: repo}
}

func (s *AlunoService) FindAll() ([]models.Aluno, error) {
	return s.repo.FindAll()
}

func (s *AlunoService) FindById(id int) (*models.Aluno, error) {
	return s.repo.FindById(id)
}

func (s *AlunoService) FindByCpf(cpf string) (*models.Aluno, error) {
	return s.repo.FindByCpf(cpf)
}

func (s *AlunoService) Create(a models.Aluno) (*models.Aluno, error) {
	if err := validarAluno(&a); err != nil {
		return nil, err
	}
	return s.repo.Create(a)
}

func (s *AlunoService) Delete(id int) (*models.Aluno, error) {
	a, err := s.repo.FindById(id)
	if err != nil {
		return nil, ErrNotFound
	}

	if err := s.repo.Delete(a); err != nil {
		return nil, err
	}

	return a, nil
}

func (s *AlunoService) Update(id int, data models.Aluno) (*models.Aluno, error) {
	if err := validarAluno(&data); err != nil {
		return nil, err
	}

	a, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	a.Nome = data.Nome
	a.CPF = data.CPF
	a.RG = data.RG

	if err := s.repo.Update(a); err != nil {
		return nil, err
	}
	return a, nil
}

func validarAluno(a *models.Aluno) error {
	a.Nome = strings.TrimSpace(a.Nome)
	a.CPF = strings.TrimSpace(a.CPF)
	a.RG = strings.TrimSpace(a.RG)

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

	if a.RG == "" {
		return ErrRgObrigatorio
	}

	return nil
}
