package dto

type CreateAlunoRequest struct {
	Nome string `json:"nome" binding:"required"`
	CPF  string `json:"cpf" binding:"required,len=11,numeric"`
	RG   string `json:"rg"`
}

type UpdateAlunoRequest struct {
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

type PatchAlunoRequest struct {
	Nome *string `json:"nome,omitempty"`
	CPF  *string `json:"cpf,omitempty"`
	RG   *string `json:"rg,omitempty"`
}
