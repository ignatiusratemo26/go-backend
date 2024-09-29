package services

import (
	"errors"
	"time"
)

// Appointment represents an appointment with a start and end time.
type Appointment struct {
	ID        string
	StartTime time.Time
	EndTime   time.Time
	Details   string
}

// AppointmentService provides methods to manage appointments.
type AppointmentService struct {
	appointments map[string]Appointment
}

// NewAppointmentService creates a new AppointmentService.
func NewAppointmentService() *AppointmentService {
	return &AppointmentService{
		appointments: make(map[string]Appointment),
	}
}

// CreateAppointment creates a new appointment.
func (s *AppointmentService) CreateAppointment(id string, startTime, endTime time.Time, details string) error {
	if _, exists := s.appointments[id]; exists {
		return errors.New("appointment with this ID already exists")
	}
	s.appointments[id] = Appointment{
		ID:        id,
		StartTime: startTime,
		EndTime:   endTime,
		Details:   details,
	}
	return nil
}

// GetAppointment retrieves an appointment by ID.
func (s *AppointmentService) GetAppointment(id string) (Appointment, error) {
	appointment, exists := s.appointments[id]
	if !exists {
		return Appointment{}, errors.New("appointment not found")
	}
	return appointment, nil
}

// UpdateAppointment updates an existing appointment.
func (s *AppointmentService) UpdateAppointment(id string, startTime, endTime time.Time, details string) error {
	if _, exists := s.appointments[id]; !exists {
		return errors.New("appointment not found")
	}
	s.appointments[id] = Appointment{
		ID:        id,
		StartTime: startTime,
		EndTime:   endTime,
		Details:   details,
	}
	return nil
}

// DeleteAppointment deletes an appointment by ID.
func (s *AppointmentService) DeleteAppointment(id string) error {
	if _, exists := s.appointments[id]; !exists {
		return errors.New("appointment not found")
	}
	delete(s.appointments, id)
	return nil
}
