package repository

import (
	"context"

	"github.com/khairulharu/averincrud/domain"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatient(db *gorm.DB) domain.PatientRepository {
	return &patientRepository{
		db: db,
	}
}

func (p patientRepository) FindAll(ctx context.Context) (patients []domain.Patient, err error) {
	err = p.db.Debug().WithContext(ctx).Table("patients").Find(&patients).Order("id ASC").Error
	return
}

func (p patientRepository) FindByName(ctx context.Context, name string) (patient domain.Patient, err error) {
	err = p.db.Debug().WithContext(ctx).Table("patients").Where("name=?", name).First(&patient).Error
	return
}

func (p patientRepository) Insert(ctx context.Context, patient *domain.Patient) error {
	err := p.db.Debug().WithContext(ctx).Table("patients").Create(patient).Error
	return err
}

func (p patientRepository) Update(ctx context.Context, patient *domain.Patient) error {
	err := p.db.Debug().WithContext(ctx).Table("patients").Save(patient).Error
	return err
}

func (p patientRepository) Delete(ctx context.Context, patient domain.Patient) error {
	err := p.db.Debug().WithContext(ctx).Table("patients").Delete(&patient).Error
	return err
}
