package dto

type AlunoResponse struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
