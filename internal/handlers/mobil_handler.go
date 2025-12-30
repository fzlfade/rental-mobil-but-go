package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"rental-mobil/internal/models"
	"rental-mobil/internal/services"
)

type MobilHandler struct {
	Service *services.MobilService
}

func NewMobilHandler(service *services.MobilService) *MobilHandler {
	return &MobilHandler{Service: service}
}

func (h *MobilHandler) GetAllMobil(w http.ResponseWriter, r *http.Request) {
	mobil, err := h.Service.GetAllMobil()
	if err != nil {
		http.Error(w, "Gagal mengambil data mobil", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mobil)
}

func (h *MobilHandler) GetMobilByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Path[len("/mobil/"):]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID mobil tidak valid", http.StatusBadRequest)
		return
	}

	mobil, err := h.Service.GetMobilByID(uint(id))
	if err != nil {
		http.Error(w, "Mobil tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mobil)
}

func (h *MobilHandler) CreateMobil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var mobil models.Mobil
	err := json.NewDecoder(r.Body).Decode(&mobil)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateMobil(&mobil)
	if err != nil {
		http.Error(w, "Gagal membuat mobil", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mobil)
}

func (h *MobilHandler) UpdateMobil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Path[len("/mobil/"):]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID mobil tidak valid", http.StatusBadRequest)
		return
	}

	var mobil models.Mobil
	err = json.NewDecoder(r.Body).Decode(&mobil)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateMobil(uint(id), &mobil)
	if err != nil {
		http.Error(w, "Gagal mengupdate mobil", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mobil)
}

func (h *MobilHandler) DeleteMobil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Path[len("/mobil/"):]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID mobil tidak valid", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteMobil(uint(id))
	if err != nil {
		http.Error(w, "Gagal menghapus mobil", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mobil deleted successfully"))
}