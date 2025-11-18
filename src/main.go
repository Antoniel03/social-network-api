package main

import (
	// "context"
	"log"
	"net/http"

	"github.com/Antoniel03/social-network-api/handler"
	"github.com/Antoniel03/social-network-api/internal/db"
	"github.com/Antoniel03/social-network-api/internal/storage"
	"github.com/Antoniel03/social-network-api/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := http.NewServeMux()
	dbConnection := db.SetupDB()

	//Migrations being used
	// err = db.CreateTables(dbConnection)
	// if err != nil {
	// 	log.Fatalf("An error ocurred while trying to create the database tables: %v", err)
	// }

	store := storage.NewStorage(dbConnection)

	a := handler.Handler{Service: &service.Service{Repository: store}}

	SetupRouter(router, &a)
	log.Fatal(http.ListenAndServe(":8080", Logger{router}))
}
