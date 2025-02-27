// internal/handlers/child_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/eduardohass/kids-api/internal/services"
	"github.com/gorilla/mux"
)

type ChildHandler struct {
	childService services.ChildService
}

func NewChildHandler(childService services.ChildService) *ChildHandler {
	return &ChildHandler{
		childService: childService,
	}
}

func (h *ChildHandler) Create(w http.ResponseWriter, r *http.Request) {
	var child models.Child
	err := json.NewDecoder(r.Body).Decode(&child)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.childService.CreateChild(r.Context(), &child)
	if err != nil {
		http.Error(w, "Error creating child: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(child)
}

func (h *ChildHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	child, err := h.childService.GetChild(r.Context(), id)
	if err != nil {
		http.Error(w, "Error fetching child: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if child == nil {
		http.Error(w, "Child not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(child)
}

func (h *ChildHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var child models.Child
	err := json.NewDecoder(r.Body).Decode(&child)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	child.ID = id
	err = h.childService.UpdateChild(r.Context(), &child)
	if err != nil {
		http.Error(w, "Error updating child: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(child)
}

func (h *ChildHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.childService.DeleteChild(r.Context(), id)
	if err != nil {
		http.Error(w, "Error deleting child: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChildHandler) List(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// Filtros
	filter := make(map[string]interface{})
	if name := queryParams.Get("name"); name != "" {
		filter["nome"] = name
	}

	if group := queryParams.Get("group_id"); group != "" {
		filter["grupo_id"] = group
	}

	// Paginação
	page := 1
	if pageStr := queryParams.Get("page"); pageStr != "" {
		pageNum, err := strconv.Atoi(pageStr)
		if err == nil && pageNum > 0 {
			page = pageNum
		}
	}

	pageSize := 20
	if pageSizeStr := queryParams.Get("page_size"); pageSizeStr != "" {
		pageSizeNum, err := strconv.Atoi(pageSizeStr)
		if err == nil && pageSizeNum > 0 {
			pageSize = pageSizeNum
		}
	}

	children, err := h.childService.ListChildren(r.Context(), filter, page, pageSize)
	if err != nil {
		http.Error(w, "Error listing children: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(children)
}
