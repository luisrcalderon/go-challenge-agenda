package usecase

import (
	"context"

	agendav1 "go-challenge-agenda/gen/agenda/v1"
	"go-challenge-agenda/services/api/internal/domain"
	"go-challenge-agenda/services/api/internal/port"
)

type DoctorUsecase struct {
	agenda port.AgendaPort
}

func NewDoctorUsecase(agenda port.AgendaPort) *DoctorUsecase {
	return &DoctorUsecase{agenda: agenda}
}

func (u *DoctorUsecase) List(ctx context.Context) ([]domain.DoctorResponse, error) {
	resp, err := u.agenda.ListDoctors(ctx, &agendav1.ListDoctorsRequest{})
	if err != nil {
		return nil, err
	}
	out := make([]domain.DoctorResponse, len(resp.Doctors))
	for i, d := range resp.Doctors {
		out[i] = doctorProtoToDTO(d)
	}
	return out, nil
}

func (u *DoctorUsecase) Get(ctx context.Context, id string) (*domain.DoctorResponse, error) {
	resp, err := u.agenda.GetDoctor(ctx, &agendav1.GetDoctorRequest{Id: id})
	if err != nil {
		return nil, err
	}
	dto := doctorProtoToDTO(resp.Doctor)
	return &dto, nil
}

func doctorProtoToDTO(d *agendav1.Doctor) domain.DoctorResponse {
	resp := domain.DoctorResponse{ID: d.Id, Name: d.Name, Specialty: d.Specialty}
	for _, wh := range d.WorkingHours {
		resp.WorkingHours = append(resp.WorkingHours, domain.WorkingHoursResp{
			Weekday: int(wh.Weekday),
			From:    wh.From,
			To:      wh.To,
		})
	}
	return resp
}
