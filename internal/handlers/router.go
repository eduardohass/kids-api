// internal/handlers/router.go
// Package handlers provides HTTP request handlers for the API endpoints.
package handlers

import (
	"github.com/eduardohass/kids-api/internal/services"
	"github.com/gorilla/mux"
)

// NewRouter creates and configures a new router with all API endpoints.
func NewRouter(
	childService services.ChildService,
	caretakerService services.CaretakerService,
	volunteerService services.VolunteerService,
	groupService services.GroupService,
) *mux.Router {
	r := mux.NewRouter()

	// Health check route (public)
	r.HandleFunc("/health", HealthHandler).Methods("GET")

	// Handlers
	childHandler := NewChildHandler(childService)
	caretakerHandler := NewCaretakerHandler(caretakerService)
	volunteerHandler := NewVolunteerHandler(volunteerService)
	groupHandler := NewGroupHandler(groupService)

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// Rotas para crianças
	api.HandleFunc("/children", childHandler.Create).Methods("POST")
	api.HandleFunc("/children", childHandler.List).Methods("GET")
	api.HandleFunc("/children/{id}", childHandler.Get).Methods("GET")
	api.HandleFunc("/children/{id}", childHandler.Update).Methods("PUT")
	api.HandleFunc("/children/{id}", childHandler.Delete).Methods("DELETE")

	// Rotas para responsáveis
	api.HandleFunc("/caretakers", caretakerHandler.Create).Methods("POST")
	api.HandleFunc("/caretakers", caretakerHandler.List).Methods("GET")
	api.HandleFunc("/caretakers/{id}", caretakerHandler.Get).Methods("GET")
	api.HandleFunc("/caretakers/{id}", caretakerHandler.Update).Methods("PUT")
	api.HandleFunc("/caretakers/{id}", caretakerHandler.Delete).Methods("DELETE")

	// Rotas para voluntários
	api.HandleFunc("/volunteers", volunteerHandler.Create).Methods("POST")
	api.HandleFunc("/volunteers", volunteerHandler.List).Methods("GET")
	api.HandleFunc("/volunteers/{id}", volunteerHandler.Get).Methods("GET")
	api.HandleFunc("/volunteers/{id}", volunteerHandler.Update).Methods("PUT")
	api.HandleFunc("/volunteers/{id}", volunteerHandler.Delete).Methods("DELETE")

	// Rotas para grupos
	api.HandleFunc("/groups", groupHandler.Create).Methods("POST")
	api.HandleFunc("/groups", groupHandler.List).Methods("GET")
	api.HandleFunc("/groups/{id}", groupHandler.Get).Methods("GET")
	api.HandleFunc("/groups/{id}", groupHandler.Update).Methods("PUT")
	api.HandleFunc("/groups/{id}", groupHandler.Delete).Methods("DELETE")

	return r
}
