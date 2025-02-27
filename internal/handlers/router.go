// internal/handlers/router.go
package handlers

import (
	"github.com/eduardohass/kids-api/internal/auth"
	"github.com/eduardohass/kids-api/internal/services"
	"github.com/gorilla/mux"
)

func NewRouter(
	auth *auth.Authenticator,
	childService services.ChildService,
	caretakerService services.CaretakerService,
	volunteerService services.VolunteerService,
	groupService services.GroupService,
) *mux.Router {
	r := mux.NewRouter()

	// Handlers
	childHandler := NewChildHandler(childService)
	caretakerHandler := NewCaretakerHandler(caretakerService)
	volunteerHandler := NewVolunteerHandler(volunteerService)
	groupHandler := NewGroupHandler(groupService)

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// Rotas públicas (se necessário)

	// Rotas autenticadas
	protected := api.PathPrefix("").Subrouter()
	protected.Use(auth.GetMiddleware())

	// Rotas para crianças
	protected.HandleFunc("/children", childHandler.Create).Methods("POST")
	protected.HandleFunc("/children", childHandler.List).Methods("GET")
	protected.HandleFunc("/children/{id}", childHandler.Get).Methods("GET")
	protected.HandleFunc("/children/{id}", childHandler.Update).Methods("PUT")
	protected.HandleFunc("/children/{id}", childHandler.Delete).Methods("DELETE")

	// Rotas para responsáveis
	protected.HandleFunc("/caretakers", caretakerHandler.Create).Methods("POST")
	protected.HandleFunc("/caretakers", caretakerHandler.List).Methods("GET")
	protected.HandleFunc("/caretakers/{id}", caretakerHandler.Get).Methods("GET")
	protected.HandleFunc("/caretakers/{id}", caretakerHandler.Update).Methods("PUT")
	protected.HandleFunc("/caretakers/{id}", caretakerHandler.Delete).Methods("DELETE")

	// Rotas para voluntários
	protected.HandleFunc("/volunteers", volunteerHandler.Create).Methods("POST")
	protected.HandleFunc("/volunteers", volunteerHandler.List).Methods("GET")
	protected.HandleFunc("/volunteers/{id}", volunteerHandler.Get).Methods("GET")
	protected.HandleFunc("/volunteers/{id}", volunteerHandler.Update).Methods("PUT")
	protected.HandleFunc("/volunteers/{id}", volunteerHandler.Delete).Methods("DELETE")

	// Rotas para grupos
	protected.HandleFunc("/groups", groupHandler.Create).Methods("POST")
	protected.HandleFunc("/groups", groupHandler.List).Methods("GET")
	protected.HandleFunc("/groups/{id}", groupHandler.Get).Methods("GET")
	protected.HandleFunc("/groups/{id}", groupHandler.Update).Methods("PUT")
	protected.HandleFunc("/groups/{id}", groupHandler.Delete).Methods("DELETE")

	return r
}
