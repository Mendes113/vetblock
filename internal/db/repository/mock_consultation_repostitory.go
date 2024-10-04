package repository

import (
	"context"
	"errors"
	"vetblock/internal/db/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// Mock do repositório
type MockConsultationRepo struct {
	mock.Mock
}

func (m *MockConsultationRepo) FindConsultationByID(ctx context.Context, id uuid.UUID) (*model.Consultation, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Consultation), args.Error(1)
}

func (m *MockConsultationRepo) SaveConsultation(ctx context.Context, consultation *model.Consultation) error {
	args := m.Called(ctx, consultation)
	return args.Error(0)
}

func (m *MockConsultationRepo) DeleteConsultation(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func MockGetVeterinaryByCRVM(crvm string) (*model.Veterinary, error) {
	if crvm == "valid-crvm" {
		return &model.Veterinary{}, nil
	}
	return nil, errors.New("veterinário não encontrado")
}

func MockGetAnimalByID(id uuid.UUID) (*model.Animal, error) {
	if id == uuid.Nil {
		return nil, errors.New("animal não encontrado")
	}
	return &model.Animal{}, nil
}
