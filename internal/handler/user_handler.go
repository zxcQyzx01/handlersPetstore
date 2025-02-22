package handler

import (
	"encoding/json"
	"handlersPetstore/internal/domain"
	"handlersPetstore/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Create user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body domain.User true "User object that needs to be added"
// @Success 200 {object} domain.User
// @Router /user [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// @Summary Create users with array
// @Description Creates list of users with given input array
// @Tags user
// @Accept json
// @Produce json
// @Param users body []domain.User true "List of user objects"
// @Success 200 {string} string "Users created successfully"
// @Router /user/createWithArray [post]
func (h *UserHandler) CreateUsersWithArrayInput(w http.ResponseWriter, r *http.Request) {
	var users []domain.User
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUsersWithArray(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Get user by username
// @Description Get user by username
// @Tags user
// @Accept json
// @Produce json
// @Param username path string true "The name that needs to be fetched"
// @Success 200 {object} domain.User
// @Router /user/{username} [get]
func (h *UserHandler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := h.service.GetUserByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// @Summary Update user
// @Description Update user by username
// @Tags user
// @Accept json
// @Produce json
// @Param username path string true "Name of user that needs to be updated"
// @Param user body domain.User true "User object that needs to be updated"
// @Success 200 {string} string "User updated successfully"
// @Router /user/{username} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateUser(username, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Delete user
// @Description Delete user by username
// @Tags user
// @Accept json
// @Produce json
// @Param username path string true "The name that needs to be deleted"
// @Success 200 {string} string "User deleted successfully"
// @Router /user/{username} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	if err := h.service.DeleteUser(username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Login user
// @Description Logs user into the system
// @Tags user
// @Accept json
// @Produce json
// @Param username query string true "The user name for login"
// @Param password query string true "The password for login"
// @Success 200 {string} string "Login successful"
// @Router /user/login [get]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	token, err := h.service.Login(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

// @Summary Logout user
// @Description Logs out current logged in user session
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} string "Logout successful"
// @Router /user/logout [get]
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Logout(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
