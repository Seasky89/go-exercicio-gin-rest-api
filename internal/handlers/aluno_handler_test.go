package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Seasky89/go-gin-rest-api/internal/domain"
	"github.com/Seasky89/go-gin-rest-api/internal/dto"
	"github.com/Seasky89/go-gin-rest-api/internal/handlers"
	"github.com/Seasky89/go-gin-rest-api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type fakeAlunoService struct {
	findAllFunc   func() ([]models.Aluno, error)
	findByIdFunc  func(int) (*models.Aluno, error)
	findByCpfFunc func(cpf string) (*models.Aluno, error)
	createFunc    func(a models.Aluno) (*models.Aluno, error)
	deleteFunc    func(id int) (*models.Aluno, error)
	updateFunc    func(id int, req dto.UpdateAlunoRequest) (*models.Aluno, error)
	patchFunc     func(id int, req dto.PatchAlunoRequest) (*models.Aluno, error)
}

func (f *fakeAlunoService) FindAll() ([]models.Aluno, error) {
	return f.findAllFunc()
}

func (f *fakeAlunoService) FindById(id int) (*models.Aluno, error) {
	return f.findByIdFunc(id)
}

func (f *fakeAlunoService) FindByCpf(cpf string) (*models.Aluno, error) {
	return f.findByCpfFunc(cpf)
}

func (f *fakeAlunoService) Create(a models.Aluno) (*models.Aluno, error) {
	return f.createFunc(a)
}

func (f *fakeAlunoService) Delete(id int) (*models.Aluno, error) {
	return f.deleteFunc(id)
}

func (f *fakeAlunoService) Update(id int, req dto.UpdateAlunoRequest) (*models.Aluno, error) {
	return f.updateFunc(id, req)
}

func (f *fakeAlunoService) Patch(id int, req dto.PatchAlunoRequest) (*models.Aluno, error) {
	return f.patchFunc(id, req)
}

func TestFindAll_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fakeService := &fakeAlunoService{
		findAllFunc: func() ([]models.Aluno, error) {
			return []models.Aluno{
				{Nome: "João", CPF: "11122233344", RG: "112223334"},
				{Nome: "Maria", CPF: "00000000000"},
			}, nil
		},
	}

	handler := handlers.NewAlunoHandler(fakeService)
	r := gin.Default()
	r.GET("/alunos", handler.FindAll)

	req, _ := http.NewRequest(http.MethodGet, "/alunos", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "João")
	assert.Contains(t, w.Body.String(), "Maria")
}

func TestFindAll_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fakeService := &fakeAlunoService{
		findAllFunc: func() ([]models.Aluno, error) {
			return nil, errors.New("erro")
		},
	}

	handler := handlers.NewAlunoHandler(fakeService)
	r := gin.Default()
	r.GET("/alunos", handler.FindAll)

	req, _ := http.NewRequest(http.MethodGet, "/alunos", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestFindById_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := handlers.NewAlunoHandler(nil)
	r := gin.Default()
	r.GET("/alunos/:id", handler.FindById)

	req, _ := http.NewRequest(http.MethodGet, "/alunos/abc", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFindByCpf_ValidCPF(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedCPF := "11122233344"

	fakeService := &fakeAlunoService{
		findByCpfFunc: func(cpf string) (*models.Aluno, error) {
			assert.Equal(t, expectedCPF, cpf)
			return &models.Aluno{
				Nome: "João",
				CPF:  cpf,
				RG:   "112223334",
			}, nil
		},
	}

	handler := handlers.NewAlunoHandler(fakeService)
	r := gin.Default()
	r.GET("/alunos/cpf/:cpf", handler.FindByCpf)

	req, _ := http.NewRequest(http.MethodGet, "/alunos/cpf/"+expectedCPF, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp models.Aluno
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "João", resp.Nome)
	assert.Equal(t, expectedCPF, resp.CPF)
}

func TestFindByCpf_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fakeService := &fakeAlunoService{
		findByCpfFunc: func(cpf string) (*models.Aluno, error) {
			return nil, domain.ErrNotFound
		},
	}

	handler := handlers.NewAlunoHandler(fakeService)
	r := gin.Default()
	r.GET("/alunos/cpf/:cpf", handler.FindByCpf)

	req, _ := http.NewRequest(http.MethodGet, "/alunos/cpf/999", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateAluno_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := handlers.NewAlunoHandler(nil)
	r := gin.Default()
	r.POST("/alunos", handler.Create)

	req, _ := http.NewRequest(http.MethodPost, "/alunos", bytes.NewBufferString("{invalid json}"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateAluno_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fakeService := &fakeAlunoService{
		createFunc: func(a models.Aluno) (*models.Aluno, error) {
			a.ID = 1
			return &a, nil
		},
	}

	handler := handlers.NewAlunoHandler(fakeService)
	r := gin.Default()
	r.POST("/alunos", handler.Create)

	body := `{"nome":"João","cpf":"11111111111","rg":"111111111"}`
	req, _ := http.NewRequest(http.MethodPost, "/alunos", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"id":1`)
}

func TestDeleteAluno_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fakeService := &fakeAlunoService{
		deleteFunc: func(id int) (*models.Aluno, error) {
			assert.Equal(t, 10, id)
			return &models.Aluno{
				Model: gorm.Model{ID: uint(10)},
				Nome:  "João",
			}, nil
		},
	}

	handler := handlers.NewAlunoHandler(fakeService)
	r := gin.Default()
	r.DELETE("/alunos/:id", handler.Delete)

	req, _ := http.NewRequest(http.MethodDelete, "/alunos/10", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteAluno_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := handlers.NewAlunoHandler(&fakeAlunoService{})
	r := gin.Default()
	r.DELETE("/alunos/:id", handler.Delete)

	req, _ := http.NewRequest(http.MethodDelete, "/alunos/abc", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateAluno_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fakeService := &fakeAlunoService{
		updateFunc: func(id int, req dto.UpdateAlunoRequest) (*models.Aluno, error) {
			assert.Equal(t, 1, id)
			assert.Equal(t, "Maria", req.Nome)
			return &models.Aluno{
				Model: gorm.Model{ID: uint(1)},
				Nome:  req.Nome,
				CPF:   "123",
				RG:    "456",
			}, nil
		},
	}

	handler := handlers.NewAlunoHandler(fakeService)
	r := gin.Default()
	r.PUT("/alunos/:id", handler.Update)

	body := `{"nome":"Maria","cpf":"123","rg":"456"}`
	req, _ := http.NewRequest(http.MethodPut, "/alunos/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateAluno_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := handlers.NewAlunoHandler(&fakeAlunoService{})
	r := gin.Default()
	r.PUT("/alunos/:id", handler.Update)

	req, _ := http.NewRequest(http.MethodPut, "/alunos/1", strings.NewReader("{invalid-json}"))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPatchAluno_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fakeService := &fakeAlunoService{
		patchFunc: func(id int, req dto.PatchAlunoRequest) (*models.Aluno, error) {
			assert.Equal(t, 2, id)
			return &models.Aluno{
				Model: gorm.Model{ID: uint(2)},
				Nome:  "Novo Nome",
			}, nil
		},
	}

	handler := handlers.NewAlunoHandler(fakeService)
	r := gin.Default()
	r.PATCH("/alunos/:id", handler.Patch)

	body := `{"nome":"Novo Nome"}`
	req, _ := http.NewRequest(http.MethodPatch, "/alunos/2", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
