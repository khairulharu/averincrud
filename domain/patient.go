package domain

import (
	"context"
	"time"

	"github.com/khairulharu/averincrud/dto"
)

type Patient struct {
	ID         int64
	Name       string
	Gender     string
	Indication string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PatientRepository interface {
	Insert(ctx context.Context, patient *Patient) error
	FindAll(ctx context.Context) ([]Patient, error)
	FindByName(ctx context.Context, name string) (Patient, error)
	Update(ctx context.Context, patient *Patient)
	Delete(ctx context.Context, name string)
}

type PatientService interface {
	GetAll(ctx context.Context) dto.Response
	StorePatient(ctx context.Context, req dto.ReqPatient) dto.Response
	UpdatePatient(ctx context.Context, req dto.ReqPatient) dto.Response
	DeletePatient(ctx context.Context, req dto.ReqPatient) dto.Response
	FindPatient(ctx context.Context, req dto.ReqPatient) dto.Response
}
