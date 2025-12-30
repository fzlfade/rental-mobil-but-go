package main

import (
	"log"
	"net/http"
	"rental-mobil/internal/database"
	"rental-mobil/internal/handlers"
	"rental-mobil/internal/repositories"
	"rental-mobil/internal/services"
)

func main() {
	database.ConnectDB()

	mobilRepo := repositories.NewMobilRepository(database.DB)
	userRepo := repositories.NewUserRepository(database.DB)
	bookingRepo := repositories.NewBookingRepository(database.DB)

	mobilSvc := services.NewMobilService(mobilRepo)
	userSvc := services.NewUserService(userRepo)
	bookingSvc := services.NewBookingService(bookingRepo, userRepo, mobilRepo)

	mobilHandler := handlers.NewMobilHandler(mobilSvc)
	userHandler := handlers.NewUserHandler(userSvc)
	bookingHandler := handlers.NewBookingHandler(bookingSvc)
	http.HandleFunc("/mobil", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			mobilHandler.GetAllMobil(w, r)
		case http.MethodPost:
			mobilHandler.CreateMobil(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/mobil/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			mobilHandler.GetMobilByID(w, r)
		case http.MethodPut:
			mobilHandler.UpdateMobil(w, r)
		case http.MethodDelete:
			mobilHandler.DeleteMobil(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetAllUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		idParam := r.URL.Path[len("/users/"):]

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

		if idParam == "" {
			http.Error(w, "ID user tidak valid", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			userHandler.GetUserByID(w, r)
		case http.MethodPut:
			userHandler.UpdateUser(w, r)
		case http.MethodDelete:
			userHandler.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			bookingHandler.CreateBooking(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/bookings/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		userPrefix := "/bookings/user/"
		if len(path) >= len(userPrefix) && path[:len(userPrefix)] == userPrefix {
			bookingHandler.GetBookingsByUserID(w, r)
			return
		}

		idParam := path[len("/bookings/"):]

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

		if idParam == "" {
			http.Error(w, "ID booking tidak valid", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			bookingHandler.GetBookingByID(w, r)
		case http.MethodPut:
			bookingHandler.UpdateBooking(w, r)
		case http.MethodDelete:
			bookingHandler.DeleteBooking(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server berjalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
