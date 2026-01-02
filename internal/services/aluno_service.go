package services

import (
	"strings"

	"github.com/Seasky89/go-gin-rest-api/internal/domain"
	"github.com/Seasky89/go-gin-rest-api/internal/dto"
	"github.com/Seasky89/go-gin-rest-api/internal/models"
	"github.com/Seasky89/go-gin-rest-api/internal/repository"
)

type AlunoService interface {
	FindAll() ([]models.Aluno, error)
	FindById(id int) (*models.Aluno, error)
	FindByCpf(cpf string) (*models.Aluno, error)
	Create(a models.Aluno) (*models.Aluno, error)
	Delete(id int) (*models.Aluno, error)
	Update(id int, req dto.UpdateAlunoRequest) (*models.Aluno, error)
	Patch(id int, req dto.PatchAlunoRequest) (*models.Aluno, error)
}

type alunoService struct {
	repo repository.AlunoRepository
}

func NewAlunoService(repo repository.AlunoRepository) AlunoService {
	return &alunoService{repo: repo}
}

func (s *alunoService) FindAll() ([]models.Aluno, error) {
	return s.repo.FindAll()
}

func (s *alunoService) FindById(id int) (*models.Aluno, error) {
	return s.repo.FindById(id)
}

func (s *alunoService) FindByCpf(cpf string) (*models.Aluno, error) {
	return s.repo.FindByCpf(cpf)
}

func (s *alunoService) Create(a models.Aluno) (*models.Aluno, error) {
	a.Nome = strings.TrimSpace(a.Nome)
	a.CPF = strings.TrimSpace(a.CPF)
	a.RG = strings.TrimSpace(a.RG)

	if err := domain.ValidarAluno(&a); err != nil {
		return nil, err
	}

	return s.repo.Create(a)
}

func (s *alunoService) Delete(id int) (*models.Aluno, error) {
	return s.repo.DeleteById(id)
}

func (s *alunoService) Update(id int, req dto.UpdateAlunoRequest) (*models.Aluno, error) {
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

func (s *alunoService) Patch(id int, req dto.PatchAlunoRequest) (*models.Aluno, error) {
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
		rg := strings.TrimSpace(*req.RG)
		if rg == "" {
			return nil, domain.ErrRgObrigatorio
		}
		aluno.RG = rg
	}

	if err := domain.ValidarAluno(aluno); err != nil {
		return nil, err
	}

	if err := s.repo.Update(aluno); err != nil {
		return nil, err
	}
	return aluno, nil
}
