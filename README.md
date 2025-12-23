# 📚 API de Alunos (Go)

API REST desenvolvida em Go para gerenciamento de alunos, seguindo boas práticas de organização em camadas (Handler, Service, Repository, Domain e DTOs).

O projeto utiliza **Gin** para o servidor HTTP e **GORM** para persistência de dados.

---

## 🚀 Tecnologias utilizadas

- Go
- Gin
- GORM
- PostgreSQL
- Docker / Docker Compose
- Postman (para testes das requisições)

---

## 📂 Estrutura do projeto

```text
go-exercicio-gin-rest-api
├── cmd/
│ └── api/
│ └── main.go
├── internal/
│ ├── handler/
│ ├── service/
│ ├── repository/
│ ├── domain/
│ ├── dto/
│ └── models/
├── migrations/
├── docker-compose.yml
├── go.mod
└── README.md
```

### 📌 Camadas

- **handler**: camada HTTP (Gin). Responsável por receber requests e devolver responses.
- **service**: regras de negócio.
- **repository**: acesso ao banco de dados.
- **domain**: validações e erros de domínio.
- **models**: entidades persistidas no banco.
- **dto**: objetos de entrada e saída da API.

---

## ⚙️ Como executar o projeto

### Pré-requisitos

- Go instalado
- Docker e Docker Compose

### Subindo o banco de dados

```bash
docker-compose up -d
```

Executando a aplicação

```bash
go run cmd/api/main.go
```

A API ficará disponível em:

```arduino
http://localhost:8080
```

---

## 📌 Endpoints disponíveis

### 🔹 Criar aluno
**POST** `/alunos`

```json
{
  "nome": "João da Silva",
  "cpf": "12345678901",
  "rg": "1234567"
}
```

---

### 🔹 Listar alunos

**GET** `/alunos`

---

### 🔹 Buscar aluno por ID

**GET** `/alunos/{id}`

---

### 🔹 Buscar aluno por CPF

**GET** `/alunos/cpf/{cpf}`

---

### 🔹 Atualizar aluno (PUT)

Atualização completa do recurso.

**PUT** `/alunos/{id}`

```json
{
  "nome": "João da Silva",
  "cpf": "12345678901",
  "rg": "7654321"
}
```

---

### 🔹 Atualizar aluno parcialmente (PATCH)
Atualização parcial do recurso. Apenas os campos enviados serão alterados.

**PATCH** `/alunos/{id}`

```json
{
  "nome": "João Silva"
}
```

---

### 🔹 Remover aluno

**DELETE** `/alunos/{id}`

---

## ✅ Validações

- Nome obrigatório
- CPF obrigatório, numérico e com 11 dígitos
- Validações centralizadas na camada **domain**
- PATCH permite atualização parcial com validação apenas dos campos enviados

---

## 🧪 Testes das requisições

As requisições da API foram testadas utilizando o **Postman**.

---

## 📌 Observações

- O PUT substitui completamente o recurso.
- O PATCH altera apenas os campos enviados.
- DTOs são usados para separar o contrato da API das entidades do domínio.
- O domínio não depende de framework HTTP.

---

## 📄 Licença

Este projeto é apenas para fins de estudo.
