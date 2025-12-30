package handlers

import (
	"encoding/json"
	"net/http"
	"rental-mobil/internal/models"
	"rental-mobil/internal/services"
	"strconv"
)

type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Path[len("/users/"):]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID user tidak valid", http.StatusBadRequest)
		return
	}

	var userInput UserInput
	err = json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Buat struct model.User dari input
	user := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	err = h.Service.UpdateUser(uint(id), &user)
	if err != nil {
		http.Error(w, "Gagal mengupdate user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Path[len("/users/"):]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID user tidak valid", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteUser(uint(id))
	if err != nil {
		http.Error(w, "Gagal menghapus user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	users, err := h.Service.GetAllUsers()
	if err != nil {
		http.Error(w, "Gagal mengambil data user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idParam := r.URL.Path[len("/users/"):]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID user tidak valid", http.StatusBadRequest)
		return
	}
	user, err := h.Service.GetUserByID(uint(id))
	if err != nil {
		http.Error(w, "User tidak ditemukan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var userInput UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	user := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
	}
	err = h.Service.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Encode model.User asli, field Password akan di-ignore karena json:"-"
	json.NewEncoder(w).Encode(user)
}
