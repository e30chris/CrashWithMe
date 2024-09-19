package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-username/go-web-app/internal/handlers"
	"github.com/your-username/go-web-app/internal/routes"
)

func main() {
	// Create a new handler
	handler := handlers.NewHandler()

	// Setup routes
	router := routes.SetupRoutes(handler)

	// Start the server
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}