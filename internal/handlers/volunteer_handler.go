// Package handlers provides the HTTP handlers for the volunteer operations.
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/eduardohass/kids-api/internal/services"
	"github.com/gorilla/mux"
)

type VolunteerHandler struct {
	service services.VolunteerService
}

func NewVolunteerHandler(service services.VolunteerService) *VolunteerHandler {
	return &VolunteerHandler{
		service: service,
	}
}

func (h *VolunteerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var volunteer models.Volunteer
	if err := json.NewDecoder(r.Body).Decode(&volunteer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateVolunteer(r.Context(), &volunteer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(volunteer)
}

func (h *VolunteerHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	volunteer, err := h.service.GetVolunteer(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(volunteer)
}

func (h *VolunteerHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var volunteer models.Volunteer
	if err := json.NewDecoder(r.Body).Decode(&volunteer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	volunteer.ID = id
	if err := h.service.UpdateVolunteer(r.Context(), &volunteer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(volunteer)
}

func (h *VolunteerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.DeleteVolunteer(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *VolunteerHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implementar paginação e filtros a partir dos query parameters
	page := 1
	pageSize := 10
	filter := make(map[string]interface{})

	volunteers, err := h.service.ListVolunteers(r.Context(), filter, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(volunteers)
}
