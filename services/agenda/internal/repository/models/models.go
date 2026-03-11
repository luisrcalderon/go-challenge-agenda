package models

import (
	"time"

	"go-challenge-agenda/services/agenda/internal/domain"
)

// Doctor is the GORM model for the doctors table.
type Doctor struct {
	ID           string         `gorm:"primaryKey"`
	Name         string         `gorm:"not null"`
	Specialty    string         `gorm:"not null"`
	WorkingHours []WorkingHours `gorm:"foreignKey:DoctorID"`
}

// WorkingHours is the GORM model for working_hours table.
type WorkingHours struct {
	ID       string `gorm:"primaryKey"`
	DoctorID string `gorm:"not null;index"`
	Weekday  int    `gorm:"not null"`
	FromTime string `gorm:"column:from_time;not null"`
	ToTime   string `gorm:"column:to_time;not null"`
}

// Patient is the GORM model for the patients table.
type Patient struct {
	ID    string `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Phone string `gorm:"not null;uniqueIndex"`
	Email string `gorm:"not null"`
}

// Reservation is the GORM model for the reservations table.
type Reservation struct {
	ID        string    `gorm:"primaryKey"`
	DoctorID  string    `gorm:"not null;index"`
	PatientID string    `gorm:"not null;index"`
	StartsAt  time.Time `gorm:"not null"`
	EndsAt    time.Time `gorm:"not null"`
	Type      int       `gorm:"not null;default:0"`
	Status    int       `gorm:"not null;default:1"`
}

// BlockedSlot is the GORM model for the blocked_slots table.
type BlockedSlot struct {
	ID              string     `gorm:"primaryKey"`
	DoctorID        string     `gorm:"not null;index"`
	StartsAt        time.Time  `gorm:"not null"`
	EndsAt          time.Time  `gorm:"not null"`
	Reason          string     `gorm:"default:''"`
	RecurrenceType  int        `gorm:"not null;default:0"`
	RecurrenceUntil *time.Time `gorm:"default:null"`
}

// ─── Mappers ──────────────────────────────────────────────────────────────────

func DoctorFromModel(m *Doctor) *domain.Doctor {
	d := &domain.Doctor{ID: m.ID, Name: m.Name, Specialty: m.Specialty}
	for _, wh := range m.WorkingHours {
		d.WorkingHours = append(d.WorkingHours, domain.WorkingHours{
			Weekday: domain.Weekday(wh.Weekday),
			From:    wh.FromTime,
			To:      wh.ToTime,
		})
	}
	return d
}

func DoctorToModel(d *domain.Doctor) *Doctor {
	m := &Doctor{ID: d.ID, Name: d.Name, Specialty: d.Specialty}
	for _, wh := range d.WorkingHours {
		m.WorkingHours = append(m.WorkingHours, WorkingHours{
			DoctorID: d.ID,
			Weekday:  int(wh.Weekday),
			FromTime: wh.From,
			ToTime:   wh.To,
		})
	}
	return m
}

func PatientFromModel(m *Patient) *domain.Patient {
	return &domain.Patient{ID: m.ID, Name: m.Name, Phone: m.Phone, Email: m.Email}
}

func PatientToModel(p *domain.Patient) *Patient {
	return &Patient{ID: p.ID, Name: p.Name, Phone: p.Phone, Email: p.Email}
}

func ReservationFromModel(m *Reservation) *domain.Reservation {
	return &domain.Reservation{
		ID:        m.ID,
		DoctorID:  m.DoctorID,
		PatientID: m.PatientID,
		StartsAt:  m.StartsAt,
		EndsAt:    m.EndsAt,
		Type:      domain.ReservationType(m.Type),
		Status:    domain.ReservationStatus(m.Status),
	}
}

func ReservationToModel(r *domain.Reservation) *Reservation {
	return &Reservation{
		ID:        r.ID,
		DoctorID:  r.DoctorID,
		PatientID: r.PatientID,
		StartsAt:  r.StartsAt.UTC(),
		EndsAt:    r.EndsAt.UTC(),
		Type:      int(r.Type),
		Status:    int(r.Status),
	}
}

func BlockedSlotFromModel(m *BlockedSlot) *domain.BlockedSlot {
	b := &domain.BlockedSlot{
		ID:             m.ID,
		DoctorID:       m.DoctorID,
		StartsAt:       m.StartsAt,
		EndsAt:         m.EndsAt,
		Reason:         m.Reason,
		RecurrenceType: domain.RecurrenceType(m.RecurrenceType),
	}
	if m.RecurrenceUntil != nil {
		t := *m.RecurrenceUntil
		b.RecurrenceUntil = &t
	}
	return b
}

func BlockedSlotToModel(b *domain.BlockedSlot) *BlockedSlot {
	m := &BlockedSlot{
		ID:             b.ID,
		DoctorID:       b.DoctorID,
		StartsAt:       b.StartsAt.UTC(),
		EndsAt:         b.EndsAt.UTC(),
		Reason:         b.Reason,
		RecurrenceType: int(b.RecurrenceType),
	}
	if b.RecurrenceUntil != nil {
		t := b.RecurrenceUntil.UTC()
		m.RecurrenceUntil = &t
	}
	return m
}
