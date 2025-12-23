package repository

import (
	"github.com/Seasky89/go-gin-rest-api/internal/models"
	"gorm.io/gorm"
)

type AlunoRepository struct {
	db *gorm.DB
}

func NewAlunoRepository(db *gorm.DB) *AlunoRepository {
	return &AlunoRepository{db: db}
}

func (r *AlunoRepository) FindAll() ([]models.Aluno, error) {
	var a []models.Aluno
	err := r.db.Find(&a).Error
	return a, err
}

func (r *AlunoRepository) FindById(id int) (*models.Aluno, error) {
	var a models.Aluno
	err := r.db.First(&a, id).Error
	return &a, err
}

func (r *AlunoRepository) FindByCpf(cpf string) (*models.Aluno, error) {
	var a models.Aluno
	err := r.db.Where(&models.Aluno{CPF: cpf}).First(&a).Error
	return &a, err
}

func (r *AlunoRepository) Create(a models.Aluno) (*models.Aluno, error) {
	err := r.db.Create(&a).Error
	return &a, err
}

func (r *AlunoRepository) Delete(a *models.Aluno) error {
	return r.db.Delete(a).Error
}

func (r *AlunoRepository) Update(a *models.Aluno) error {
	return r.db.Model(&a).UpdateColumns(a).Error
}
