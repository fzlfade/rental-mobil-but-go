package handlers

import (
	"encoding/json"
	"net/http"
	"rental-mobil/internal/models"
	"rental-mobil/internal/services"
	"strconv"
)

type BookingHandler struct {
	Service *services.BookingService
}

func NewBookingHandler(service *services.BookingService) *BookingHandler {
	return &BookingHandler{Service: service}
}

func (h *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Path[len("/bookings/"):]
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID booking tidak valid", http.StatusBadRequest)
		return
	}

	booking, err := h.Service.GetBookingByID(uint(id))
	if err != nil {
		http.Error(w, "Booking tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) GetBookingsByUserID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	basePath := "/bookings/user/"
	path := r.URL.Path

	if len(path) <= len(basePath) || path[:len(basePath)] != basePath {
		http.Error(w, "Path tidak valid", http.StatusBadRequest)
		return
	}

	idParam := path[len(basePath):]
	endIndex := -1
	for i, char := range idParam {
		if string(char) == "/" {
			endIndex = i
			break
		}
	}
	if endIndex != -1 {
		idParam = idParam[:endIndex]
	}

	userID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID user tidak valid", http.StatusBadRequest)
		return
	}

	bookings, err := h.Service.GetBookingsByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var booking models.Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = h.Service.CreateBooking(&booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Path[len("/bookings/"):]
	endIndex := -1
	for i, char := range idParam {
		if string(char) == "/" {
			endIndex = i
			break
		}
	}
	if endIndex != -1 {
		idParam = idParam[:endIndex]
	}

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID booking tidak valid", http.StatusBadRequest)
		return
	}

	var booking models.Booking
	err = json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateBooking(uint(id), &booking)
	if err != nil {
		http.Error(w, "Gagal mengupdate booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Path[len("/bookings/"):]
	endIndex := -1
	for i, char := range idParam {
		if string(char) == "/" {
			endIndex = i
			break
		}
	}
	if endIndex != -1 {
		idParam = idParam[:endIndex]
	}

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || idParam == "" {
		http.Error(w, "ID booking tidak valid", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteBooking(uint(id))
	if err != nil {
		http.Error(w, "Gagal menghapus booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Booking deleted successfully"))
}
