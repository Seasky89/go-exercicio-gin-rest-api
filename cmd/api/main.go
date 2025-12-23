package main

import (
	"github.com/Seasky89/go-gin-rest-api/internal/database"
	"github.com/Seasky89/go-gin-rest-api/internal/handlers"
	"github.com/Seasky89/go-gin-rest-api/internal/repository"
	"github.com/Seasky89/go-gin-rest-api/internal/services"
	"github.com/Seasky89/go-gin-rest-api/routes"
)

func main() {
	db := database.ConnectDB()

	repo := repository.NewAlunoRepository(db)
	service := services.NewAlunoService(repo)
	handler := handlers.NewAlunoHandler(service)

	routes.HandleRequests(handler)
}
