package handler

import (
	"encoding/json"
	"handlersPetstore/internal/domain"
	"handlersPetstore/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PetHandler struct {
	service *service.PetService
}

func NewPetHandler(service *service.PetService) *PetHandler {
	return &PetHandler{service: service}
}

// @Summary Create a pet
// @Description Add a new pet to the store
// @Tags pet
// @Accept json
// @Produce json
// @Param pet body domain.Pet true "Pet object that needs to be added"
// @Success 200 {object} domain.Pet
// @Router /pet [post]
// @Security OAuth2Application[write]
func (h *PetHandler) CreatePet(w http.ResponseWriter, r *http.Request) {
	var pet domain.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreatePet(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Get pet by ID
// @Description Returns a single pet
// @Tags pet
// @Accept json
// @Produce json
// @Param petId path int true "ID of pet to return"
// @Success 200 {object} domain.Pet
// @Router /pet/{petId} [get]
// @Security ApiKeyAuth
func (h *PetHandler) GetPetByID(w http.ResponseWriter, r *http.Request) {
	petID := chi.URLParam(r, "petId")
	id, err := strconv.ParseInt(petID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid pet ID", http.StatusBadRequest)
		return
	}

	pet, err := h.service.GetPetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
}

func (h *PetHandler) FindPetsByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query()["status"]
	pets, err := h.service.FindPetsByStatus(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func (h *PetHandler) UpdatePet(w http.ResponseWriter, r *http.Request) {
	var pet domain.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdatePet(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PetHandler) DeletePet(w http.ResponseWriter, r *http.Request) {
	petID := chi.URLParam(r, "petId")
	id, err := strconv.ParseInt(petID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid pet ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeletePet(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
