package domain

// Request DTOs

type CreateReservationRequest struct {
	DoctorID     string `json:"doctor_id" binding:"required"`
	StartsAt     string `json:"starts_at" binding:"required"` // RFC3339
	Type         string `json:"type" binding:"required"`       // "first_visit" | "follow_up"
	PatientID    string `json:"patient_id"`                    // set if returning patient
	PatientName  string `json:"patient_name"`
	PatientPhone string `json:"patient_phone"`
	PatientEmail string `json:"patient_email"`
}

type GetAvailabilityRequest struct {
	Date string `form:"date" binding:"required"` // "2024-03-15"
	Type string `form:"type"`                    // "first_visit" | "follow_up"
}

// Response DTOs

type AvailableSlot struct {
	StartsAt string `json:"starts_at"`
	EndsAt   string `json:"ends_at"`
}

type TimeRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type AvailabilityResponse struct {
	Slots      []AvailableSlot `json:"slots"`
	FreeRanges []TimeRange     `json:"free_ranges"`
}

type ReservationResponse struct {
	ID        string `json:"id"`
	DoctorID  string `json:"doctor_id"`
	PatientID string `json:"patient_id"`
	StartsAt  string `json:"starts_at"`
	EndsAt    string `json:"ends_at"`
	Type      string `json:"type"`
	Status    string `json:"status"`
}

type DoctorResponse struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Specialty    string             `json:"specialty"`
	WorkingHours []WorkingHoursResp `json:"working_hours"`
}

type WorkingHoursResp struct {
	Weekday int    `json:"weekday"`
	From    string `json:"from"`
	To      string `json:"to"`
}

// User DTOs

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"  binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
