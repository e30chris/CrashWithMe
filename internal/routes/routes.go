package routes

import (
	"net/http"

	"github.com/your-username/go-web-app/internal/handlers"
)

// SetupRoutes sets up the routes for the web application.
func SetupRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/products", handlers.ProductsHandler)
}