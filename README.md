# 📚 API de Alunos (Go)

Projeto desenvolvido como exercício prático para estudo de **Go (Golang)** utilizando o framework **Gin**, aplicando boas práticas de organização de código, separação de responsabilidades e testes.

O projeto expõe uma API REST para gerenciamento de alunos e também possui rotas para páginas HTML (server-side rendering).


---

## 🛠️ Tecnologias utilizadas

- **Go**
- **Gin Gonic**
- **GORM**
- **PostgreSQL**
- **HTML Templates**
- **Docker / Docker Compose**
- **Testify** (testes)
- **Postman** (para testes das requisições)

---

## 📂 Estrutura do projeto

```text
go-exercicio-gin-rest-api
├── cmd/
│  └── api/
│      └── main.go
├── internal/
│ ├── database/        # Conexão com o banco de dados
│ ├── handler/         # Handlers HTTP (API e Pages)
│ ├── service/         # Camada de serviços (casos de uso)
│ ├── repository/      # Camada de persistência
│ ├── domain/          # Regras de negócio e validações
│ ├── dto/             # DTOs de entrada e saída da API
│ └── models/          # Models utilizados pelo GORM
├── routes/
│ ├── api_routes.go    # Rotas da API
│ ├── page_routes.go   # Rotas de páginas HTML
│ └── router.go        # Setup do router Gin
├── assets/            # Pasta com arquivos CSS
├── templates/         # Pasta com arquivos das páginas HTML
├── docker-compose.yml
├── go.mod
└── README.md
```

### 📌 Arquitetura

O projeto segue uma separação em camadas inspirada em Clean Architecture

- **Handlers**: camada HTTP (Gin). Responsável por receber requests e devolver responses.
- **Services**: regras de negócio.
- **Repository**: acesso ao banco de dados.
- **Domain**: validações e erros de domínio.
- **Models**: entidades persistidas no banco.
- **DTOs**: objetos de entrada e saída da API.

---

## ⚙️ Como executar o projeto

### Pré-requisitos

- Go instalado
- Docker e Docker Compose

### Subindo o banco de dados

```bash
docker-compose up -d
```

### Executando a aplicação

```bash
go run cmd/api/main.go
```

A API ficará disponível em:

```arduino
http://localhost:8080
```

---

## 🔌 Endpoints disponíveis (API)

### 🔹 Listar alunos

**GET** `/alunos`

---

### 🔹 Buscar aluno por ID

**GET** `/alunos/{id}`

---

### 🔹 Buscar aluno por CPF

**GET** `/alunos/cpf/{cpf}`

---

### 🔹 Criar novo aluno

**POST** `/alunos`

```json
{
  "nome": "João da Silva",
  "cpf": "12345678901",
  "rg": "1234567"
}
```

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

## 📄 Rotas de páginas (HTML)

| Método | Rota        | Descrição                    |
|--------|-------------|------------------------------|
| GET    | `/index`    | Página de listagem de alunos |

---

## ✅ Validações

- Nome obrigatório
- CPF obrigatório, numérico e com 11 dígitos
- Validações centralizadas na camada **domain**
- PATCH permite atualização parcial com validação apenas dos campos enviados

---

## 🧪 Testes

O projeto possui testes automatizados para:
- Services
- Handlers
- Rotas

Para executar os testes:

```bash
go test ./...
```

Nos testes de rotas, são utilizados handlers fake, evitando dependências de banco de dados ou templates HTML.

---

## 🔍 Testes manuais

As requisições da API foram testadas utilizando o **Postman**.

---

## 📌 Observações

- O PUT substitui completamente o recurso.
- O PATCH altera apenas os campos enviados.
- DTOs são usados para separar o contrato da API das entidades do domínio.
- O projeto tem fins educacionais, mas segue padrões próximos aos utilizados em projetos reais.

---

## 📄 Licença

Este projeto é apenas para fins de estudo.
