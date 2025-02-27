package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/eduardohass/kids-api/internal/services"
	"github.com/gorilla/mux"
)

type CaretakerHandler struct {
	service services.CaretakerService
}

func NewCaretakerHandler(service services.CaretakerService) *CaretakerHandler {
	return &CaretakerHandler{
		service: service,
	}
}

func (h *CaretakerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var caretaker models.Caretaker
	if err := json.NewDecoder(r.Body).Decode(&caretaker); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCaretaker(r.Context(), &caretaker); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(caretaker)
}

func (h *CaretakerHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	caretaker, err := h.service.GetCaretaker(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(caretaker)
}

func (h *CaretakerHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var caretaker models.Caretaker
	if err := json.NewDecoder(r.Body).Decode(&caretaker); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	caretaker.ID = id
	if err := h.service.UpdateCaretaker(r.Context(), &caretaker); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(caretaker)
}

func (h *CaretakerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.DeleteCaretaker(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CaretakerHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implementar paginação e filtros a partir dos query parameters
	page := 1
	pageSize := 10
	filter := make(map[string]interface{})

	caretakers, err := h.service.ListCaretakers(r.Context(), filter, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(caretakers)
}
