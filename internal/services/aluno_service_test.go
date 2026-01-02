package services_test

import (
	"errors"
	"testing"

	"github.com/Seasky89/go-gin-rest-api/internal/models"
	"github.com/Seasky89/go-gin-rest-api/internal/services"
	"github.com/stretchr/testify/assert"
)

type fakeAlunoRepository struct {
	createFunc func(a models.Aluno) (*models.Aluno, error)
}

func (f *fakeAlunoRepository) FindAll() ([]models.Aluno, error) {
	return nil, nil
}

func (f *fakeAlunoRepository) FindById(id int) (*models.Aluno, error) {
	return nil, nil
}

func (f *fakeAlunoRepository) FindByCpf(cpf string) (*models.Aluno, error) {
	return nil, nil
}

func (f *fakeAlunoRepository) Create(a models.Aluno) (*models.Aluno, error) {
	return f.createFunc(a)
}

func (f *fakeAlunoRepository) Delete(a *models.Aluno) error {
	return nil
}

func (f *fakeAlunoRepository) DeleteById(id int) (*models.Aluno, error) {
	return nil, nil
}

func (f *fakeAlunoRepository) Update(a *models.Aluno) error {
	return nil
}

func TestCreateAluno_TrimFields(t *testing.T) {
	fakeRepo := &fakeAlunoRepository{
		createFunc: func(a models.Aluno) (*models.Aluno, error) {
			return &a, nil
		},
	}

	services := services.NewAlunoService(fakeRepo)

	aluno := models.Aluno{
		Nome: "   João   ",
		CPF:  " 12345678901 ",
		RG:   " 112223334 ",
	}

	created, _ := services.Create(aluno)

	assert.Equal(t, "João", created.Nome)
	assert.Equal(t, "12345678901", created.CPF)
	assert.Equal(t, "112223334", created.RG)
}

func TestCreateAluno_InvalidCPF(t *testing.T) {
	fakeRepo := &fakeAlunoRepository{
		createFunc: func(a models.Aluno) (*models.Aluno, error) {
			t.Fatal("repo.Create não deveria ser chamado")
			return nil, nil
		},
	}

	service := services.NewAlunoService(fakeRepo)

	aluno := models.Aluno{
		Nome: "João",
		CPF:  "",
	}

	created, err := service.Create(aluno)

	assert.Nil(t, created)
	assert.Error(t, err)
}

func TestCreateAluno_RepoError(t *testing.T) {
	expectedErr := errors.New("erro no banco")

	fakeRepo := &fakeAlunoRepository{
		createFunc: func(a models.Aluno) (*models.Aluno, error) {
			return nil, expectedErr
		},
	}

	service := services.NewAlunoService(fakeRepo)

	aluno := models.Aluno{
		Nome: "João",
		CPF:  "12345678901",
	}

	created, err := service.Create(aluno)

	assert.Nil(t, created)
	assert.Equal(t, expectedErr, err)
}
