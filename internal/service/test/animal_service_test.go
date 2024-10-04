package service_test

import (
	"errors"
	"testing"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"
	"vetblock/internal/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Test for AddAnimal
func TestAddAnimal(t *testing.T) {
	animal := model.Animal{
		Name:    "Rex",
		Species: "Dog",
	}

	t.Run("should add valid animal", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindByUniqueAttributes", animal).Return(nil, nil)
		mockRepo.On("SaveAnimal", &animal).Return(nil)

		err := animalService.AddAnimal(animal)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t) // Ensure mock expectations are met
	})

	t.Run("should not add animal if already exists", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindByUniqueAttributes", animal).Return(&animal, nil)

		err := animalService.AddAnimal(animal)
		assert.EqualError(t, err, "animal already exists")
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindByUniqueAttributes", animal).Return(nil, errors.New("repository error"))

		err := animalService.AddAnimal(animal)
		assert.EqualError(t, err, "error checking for existing animal: repository error")
		mockRepo.AssertExpectations(t)
	})
}

// Test for GetAnimalByID
func TestGetAnimalByID(t *testing.T) {
	animalID := uuid.New()
	animal := &model.Animal{
		ID:      animalID,
		Name:    "Rex",
		Species: "Dog",
	}

	t.Run("should return animal when found", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindAnimalByID", animalID).Return(animal, nil)

		result, err := animalService.GetAnimalByID(animalID)
		assert.NoError(t, err)
		assert.Equal(t, animal, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if animal not found", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindAnimalByID", animalID).Return(nil, errors.New("animal not found"))

		result, err := animalService.GetAnimalByID(animalID)
		assert.Nil(t, result)
		assert.EqualError(t, err, "error fetching animal by ID: animal not found")
		mockRepo.AssertExpectations(t)
	})
}

// Test for UpdateAnimal
func TestUpdateAnimal(t *testing.T) {
	animalID := uuid.New()
	existingAnimal := &model.Animal{
		ID:      animalID,
		Name:    "Rex",
		Species: "Dog",
	}

	updatedAnimal := model.Animal{
		Name:    "Max",
		Species: "Dog",
		Breed:   "Golden Retriever",
		Age:     3,
	}

	t.Run("should update animal", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindAnimalByID", animalID).Return(existingAnimal, nil)
		mockRepo.On("SaveAnimal", existingAnimal).Return(nil)

		err := animalService.UpdateAnimal(animalID, updatedAnimal)
		assert.NoError(t, err)
		assert.Equal(t, updatedAnimal.Name, existingAnimal.Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if animal not found", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindAnimalByID", animalID).Return(nil, errors.New("animal not found"))

		err := animalService.UpdateAnimal(animalID, updatedAnimal)
		assert.EqualError(t, err, "error finding animal to update: animal not found")
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if save fails", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindAnimalByID", animalID).Return(existingAnimal, nil)
		mockRepo.On("SaveAnimal", existingAnimal).Return(errors.New("save failed"))

		err := animalService.UpdateAnimal(animalID, updatedAnimal)
		assert.EqualError(t, err, "error updating animal: save failed")
		mockRepo.AssertExpectations(t)
	})
}

// Test for DeleteAnimal
func TestDeleteAnimal(t *testing.T) {
	animalID := uuid.New()

	t.Run("should delete animal", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("DeleteAnimal", animalID).Return("Animal excluído com sucesso", nil)

		msg, err := animalService.DeleteAnimal(animalID)
		assert.NoError(t, err)
		assert.Equal(t, "Animal excluído com sucesso", msg)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if deletion fails", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("DeleteAnimal", animalID).Return("", errors.New("delete failed"))

		msg, err := animalService.DeleteAnimal(animalID)
		assert.EqualError(t, err, "delete failed")
		assert.Empty(t, msg)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if animal not found", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("DeleteAnimal", animalID).Return("Animal not found", errors.New("animal not found"))

		msg, err := animalService.DeleteAnimal(animalID)
		assert.EqualError(t, err, "animal not found")
		assert.Equal(t, "Animal not found", msg)
		mockRepo.AssertExpectations(t)
	})
}

// Test for GetAllAnimals
func TestGetAllAnimals(t *testing.T) {
	animals := []model.Animal{
		{Name: "Rex", Species: "Dog"},
		{Name: "Max", Species: "Cat"},
	}

	t.Run("should return all animals", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindAllAnimals").Return(animals, nil)

		result, err := animalService.GetAllAnimals()
		assert.NoError(t, err)
		assert.Equal(t, animals, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return empty list when no animals found", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindAllAnimals").Return([]model.Animal{}, nil)

		result, err := animalService.GetAllAnimals()
		assert.NoError(t, err)
		assert.Empty(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when fetching animals fails", func(t *testing.T) {
		mockRepo := new(repository.MockAnimalRepository) // Fresh mock per test
		animalService := service.NewAnimalService(mockRepo)

		mockRepo.On("FindAllAnimals").Return(nil, errors.New("fetch failed"))

		result, err := animalService.GetAllAnimals()
		assert.Nil(t, result)
		assert.EqualError(t, err, "error retrieving animals: fetch failed")
		mockRepo.AssertExpectations(t)
	})
}
