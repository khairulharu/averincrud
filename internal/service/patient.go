package service

import (
	"context"
	"time"

	"github.com/khairulharu/averincrud/domain"
	"github.com/khairulharu/averincrud/dto"
)

type patientService struct {
	patientRepository domain.PatientRepository
}

func NewPatient(patientRepository domain.PatientRepository) domain.PatientService {
	return &patientService{
		patientRepository: patientRepository,
	}
}

func (p patientService) DeletePatient(ctx context.Context, req dto.ReqPatient) dto.Response {
	patient, err := p.patientRepository.FindByName(ctx, req.Name)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Message: "Failed",
			Error:   err.Error(),
			Data:    nil,
		}
	}
	if patient == (domain.Patient{}) {
		return dto.Response{
			Code:    "404",
			Message: "Patient Not Found",
		}
	}
	if err := p.patientRepository.Delete(ctx, patient); err != nil {
		return dto.Response{
			Code:    "400",
			Message: "Failed",
			Error:   err.Error(),
			Data:    nil,
		}
	}

	return dto.Response{
		Code:    "200",
		Message: "APPROVE",
	}
}

func (p patientService) FindPatient(ctx context.Context, req dto.ReqPatient) dto.Response {
	patient, err := p.patientRepository.FindByName(ctx, req.Name)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Message: "Failed",
			Error:   err.Error(),
			Data:    nil,
		}
	}
	if patient == (domain.Patient{}) {
		return dto.Response{
			Code:    "404",
			Message: "Patient Not Found",
		}
	}

	return dto.Response{
		Code:    "200",
		Message: "APPROVE",
		Data: dto.ResPatient{
			ID:         patient.ID,
			Name:       patient.Name,
			Gender:     patient.Gender,
			Indication: patient.Indication,
			UpdatedAt:  patient.UpdatedAt,
			CreatedAt:  patient.CreatedAt,
		},
	}
}

func (p patientService) GetAll(ctx context.Context) dto.Response {
	var resPatients []dto.ResPatient
	patients, err := p.patientRepository.FindAll(ctx)
	if err != nil {
		return dto.Response{
			Code:    "404",
			Message: "Failed",
			Error:   err.Error(),
			Data:    nil,
		}
	}
	for _, v := range patients {
		resPatients = append(resPatients, dto.ResPatient{
			ID:         v.ID,
			Name:       v.Name,
			Gender:     v.Gender,
			Indication: v.Indication,
			UpdatedAt:  v.UpdatedAt,
			CreatedAt:  v.CreatedAt,
		})
	}
	return dto.Response{
		Code:    "200",
		Message: "APPROVE",
		Data:    resPatients,
	}

}

func (p patientService) StorePatient(ctx context.Context, req dto.ReqPatient) dto.Response {
	var patient = domain.Patient{
		Name:       req.Name,
		Gender:     req.Gender,
		Indication: req.Indication,
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	}

	if err := p.patientRepository.Insert(ctx, &patient); err != nil {
		return dto.Response{
			Code:    "400",
			Message: "Failed",
			Error:   err.Error(),
			Data:    nil,
		}
	}

	return dto.Response{
		Code:    "200",
		Message: "APPROVE",
	}

}

func (p patientService) UpdatePatient(ctx context.Context, name string, req dto.ReqPatient) dto.Response {
	patient, err := p.patientRepository.FindByName(ctx, name)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Message: "Failed",
			Error:   err.Error(),
			Data:    nil,
		}
	}
	if patient == (domain.Patient{}) {
		return dto.Response{
			Code:    "404",
			Message: "Patient Not Found",
		}
	}
	var savePatient = domain.Patient{
		ID:         patient.ID,
		Name:       req.Name,
		Gender:     req.Gender,
		Indication: req.Indication,
		UpdatedAt:  time.Now(),
		CreatedAt:  patient.CreatedAt,
	}

	if err := p.patientRepository.Update(ctx, &savePatient); err != nil {
		return dto.Response{
			Code:    "401",
			Message: "Failed",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    "200",
		Message: "APPROVE",
		Data: dto.ResPatient{
			ID:         savePatient.ID,
			Name:       savePatient.Name,
			Gender:     savePatient.Gender,
			Indication: savePatient.Indication,
			UpdatedAt:  savePatient.UpdatedAt,
			CreatedAt:  savePatient.CreatedAt,
		},
	}

}
