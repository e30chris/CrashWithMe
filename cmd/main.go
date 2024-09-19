package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-web-app/internal/db"
	"go-web-app/internal/handlers"
	"go-web-app/internal/routes"
)

func main() {
	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Create tables
	if err := database.CreateTables(); err != nil {
		log.Fatalf("Could not create tables: %v", err)
	}

	handler := handlers.NewHandler(database)
	router := routes.NewRouter(handler)

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
