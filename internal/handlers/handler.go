package handlers

import (
	"fmt"
	"net/http"

	"github.com/your-username/go-web-app/internal/models"
)

type Handler struct {
	Model *models.Model
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET request
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	// Handle POST request
}

func (h *Handler) PutHandler(w http.ResponseWriter, r *http.Request) {
	// Handle PUT request
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Handle DELETE request
}

func (h *Handler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 Not Found")
}