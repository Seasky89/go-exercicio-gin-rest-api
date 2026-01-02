package main

import (
	"log"

	"github.com/Seasky89/go-gin-rest-api/internal/database"
	"github.com/Seasky89/go-gin-rest-api/internal/handlers"
	"github.com/Seasky89/go-gin-rest-api/internal/repository"
	"github.com/Seasky89/go-gin-rest-api/internal/services"
	"github.com/Seasky89/go-gin-rest-api/routes"
)

func main() {
	db, err := database.ConnectDB(true)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewAlunoRepository(db)
	service := services.NewAlunoService(repo)
	handler := handlers.NewAlunoHandler(service)

	r := routes.SetupRouter()

	routes.RegisterAPIRoutes(r, handler)
	routes.RegisterPageRoutes(r, handler)

	r.Run()
}
