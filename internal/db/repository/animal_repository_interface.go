package repository

import (
    "github.com/google/uuid"
    "vetblock/internal/db/model"
)

// AnimalRepositoryInterface define os métodos que o repositório de animais deve implementar
type AnimalRepositoryInterface interface {
    FindByUniqueAttributes(animal model.Animal) (*model.Animal, error)
    SaveAnimal(animal *model.Animal) error
    FindAnimalByID(id uuid.UUID) (*model.Animal, error)
    DeleteAnimal(id uuid.UUID) (string, error)
    FindAllAnimals() ([]model.Animal, error)
}
