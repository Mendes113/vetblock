package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"vetblock/internal/db/model"
)

type MockAnimalRepository struct {
	mock.Mock
}

func (m *MockAnimalRepository) FindByUniqueAttributes(animal model.Animal) (*model.Animal, error) {
	args := m.Called(animal)
	if obj, ok := args.Get(0).(*model.Animal); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAnimalRepository) SaveAnimal(animal *model.Animal) error {
	args := m.Called(animal)
	return args.Error(0)
}

func (m *MockAnimalRepository) FindAnimalByID(id uuid.UUID) (*model.Animal, error) {
	args := m.Called(id)
	if obj, ok := args.Get(0).(*model.Animal); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAnimalRepository) DeleteAnimal(id uuid.UUID) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *MockAnimalRepository) FindAllAnimals() ([]model.Animal, error) {
	args := m.Called()
	if obj, ok := args.Get(0).([]model.Animal); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
