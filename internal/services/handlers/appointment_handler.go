package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Appointment struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Time string `json:"time"`
}

var appointments []Appointment

func GetAppointments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

func GetAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range appointments {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Appointment not found", http.StatusNotFound)
}

func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var appointment Appointment
	_ = json.NewDecoder(r.Body).Decode(&appointment)
	appointments = append(appointments, appointment)
	json.NewEncoder(w).Encode(appointment)
}

func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range appointments {
		if item.ID == params["id"] {
			appointments = append(appointments[:index], appointments[index+1:]...)
			var appointment Appointment
			_ = json.NewDecoder(r.Body).Decode(&appointment)
			appointment.ID = params["id"]
			appointments = append(appointments, appointment)
			json.NewEncoder(w).Encode(appointment)
			return
		}
	}
	http.Error(w, "Appointment not found", http.StatusNotFound)
}

func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range appointments {
		if item.ID == params["id"] {
			appointments = append(appointments[:index], appointments[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(appointments)
}
