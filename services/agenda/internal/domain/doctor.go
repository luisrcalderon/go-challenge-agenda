package domain

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type WorkingHours struct {
	Weekday Weekday
	From    string // "09:00"
	To      string // "17:00"
}

type Doctor struct {
	ID           string
	Name         string
	Specialty    string
	WorkingHours []WorkingHours
}
