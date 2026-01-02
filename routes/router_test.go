package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Seasky89/go-gin-rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type fakeAlunoHandler struct{}

func (f *fakeAlunoHandler) Welcome(c *gin.Context)   { c.Status(200) }
func (f *fakeAlunoHandler) FindAll(c *gin.Context)   { c.Status(200) }
func (f *fakeAlunoHandler) FindById(c *gin.Context)  { c.Status(200) }
func (f *fakeAlunoHandler) Create(c *gin.Context)    { c.Status(200) }
func (f *fakeAlunoHandler) Delete(c *gin.Context)    { c.Status(200) }
func (f *fakeAlunoHandler) Update(c *gin.Context)    { c.Status(200) }
func (f *fakeAlunoHandler) Patch(c *gin.Context)     { c.Status(200) }
func (f *fakeAlunoHandler) FindByCpf(c *gin.Context) { c.Status(200) }

func TestAlunosRoutesExist(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := &fakeAlunoHandler{}
	r := routes.SetupRouter()
	routes.RegisterAPIRoutes(r, handler)

	tests := []struct {
		name   string
		method string
		path   string
	}{
		{"GET /:nome", http.MethodGet, "/Marcel"},
		{"GET /alunos", http.MethodGet, "/alunos"},
		{"GET /alunos/:id", http.MethodGet, "/alunos/1"},
		{"POST /alunos", http.MethodPost, "/alunos"},
		{"DELETE /alunos/:id", http.MethodDelete, "/alunos/1"},
		{"PUT /alunos/:id", http.MethodPut, "/alunos/1"},
		{"PATCH /alunos/:id", http.MethodPatch, "/alunos/1"},
		{"GET /alunos/cpf/:cpf", http.MethodGet, "/alunos/12312312311"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.NotEqual(t, http.StatusNotFound, w.Code, "rota %s (method: %s) não registrada", tt.path, tt.method)
		})
	}
}
