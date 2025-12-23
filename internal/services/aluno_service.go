package services

import (
	"strings"

	"github.com/Seasky89/go-gin-rest-api/internal/domain"
	"github.com/Seasky89/go-gin-rest-api/internal/dto"
	"github.com/Seasky89/go-gin-rest-api/internal/models"
	"github.com/Seasky89/go-gin-rest-api/internal/repository"
)

type AlunoService struct {
	repo *repository.AlunoRepository
}

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
	a.Nome = strings.TrimSpace(a.Nome)
	a.CPF = strings.TrimSpace(a.CPF)
	a.RG = strings.TrimSpace(a.RG)

	if err := domain.ValidarAluno(&a); err != nil {
		return nil, err
	}

	return s.repo.Create(a)
}

func (s *AlunoService) Delete(id int) (*models.Aluno, error) {
	return s.repo.DeleteById(id)
}

func (s *AlunoService) Update(id int, req dto.UpdateAlunoRequest) (*models.Aluno, error) {
	aluno, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	aluno.Nome = strings.TrimSpace(req.Nome)
	aluno.CPF = strings.TrimSpace(req.CPF)
	aluno.RG = strings.TrimSpace(req.RG)

	if err := domain.ValidarAluno(aluno); err != nil {
		return nil, err
	}

	if err := s.repo.Update(aluno); err != nil {
		return nil, err
	}
	return aluno, nil
}

func (s *AlunoService) Patch(id int, req dto.PatchAlunoRequest) (*models.Aluno, error) {
	aluno, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if req.Nome != nil {
		nome := strings.TrimSpace(*req.Nome)
		if nome == "" {
			return nil, domain.ErrNomeObrigatorio
		}
		aluno.Nome = nome
	}

	if req.CPF != nil {
		cpf := strings.TrimSpace(*req.CPF)
		if cpf == "" {
			return nil, domain.ErrCpfObrigatorio
		}
		aluno.CPF = cpf
	}

	if req.RG != nil {
		aluno.RG = strings.TrimSpace(*req.RG)
	}

	if err := domain.ValidarAluno(aluno); err != nil {
		return nil, err
	}

	if err := s.repo.Update(aluno); err != nil {
		return nil, err
	}
	return aluno, nil
}
