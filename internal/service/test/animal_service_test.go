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

// Teste para AddAnimal
func TestAddAnimal(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo) // Injetando o mock no serviço

	animal := model.Animal{
		Name:    "Rex",
		Species: "Dog",
	}

	t.Run("should add valid animal", func(t *testing.T) {
		mockRepo.On("FindByUniqueAttributes", animal).Return(nil, nil)
		mockRepo.On("SaveAnimal", &animal).Return(nil)

		err := service.AddAnimal(animal)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)


	t.Run("should not add animal if already exists", func(t *testing.T) {
		mockRepo.On("FindByUniqueAttributes", animal).Return(&animal, nil)

		err := service.AddAnimal(animal)
		assert.EqualError(t, err, "animal already exists")
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)

	
	t.Run("should return error if repository fails", func(t *testing.T) {
		mockRepo.On("FindByUniqueAttributes", animal).Return(nil, errors.New("repository error"))

		err := service.AddAnimal(animal)
		assert.EqualError(t, err, "repository error")
		mockRepo.AssertExpectations(t)
	})
}

// Teste para GetAnimalByID
func TestGetAnimalByID(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)

	animalID := uuid.New()
	animal := &model.Animal{
		ID:      animalID,
		Name:    "Rex",
		Species: "Dog",
	}

	t.Run("should return animal when found", func(t *testing.T) {
		mockRepo.On("FindAnimalByID", animalID).Return(animal, nil)

		result, err := service.GetAnimalByID(animalID)
		assert.NoError(t, err)
		assert.Equal(t, animal, result)
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)


	t.Run("should return error if animal not found", func(t *testing.T) {
		mockRepo.On("FindAnimalByID", animalID).Return(nil, errors.New("animal not found"))

		result, err := service.GetAnimalByID(animalID)
		assert.Nil(t, result)
		assert.EqualError(t, err, "animal not found")
		mockRepo.AssertExpectations(t)
	})
}

// Teste para UpdateAnimal
func TestUpdateAnimal(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)

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
		mockRepo.On("FindAnimalByID", animalID).Return(existingAnimal, nil)
		mockRepo.On("SaveAnimal", existingAnimal).Return(nil)

		err := service.UpdateAnimal(animalID, updatedAnimal)
		assert.NoError(t, err)
		assert.Equal(t, updatedAnimal.Name, existingAnimal.Name)
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)

	t.Run("should return error if animal not found", func(t *testing.T) {
		mockRepo.On("FindAnimalByID", animalID).Return(nil, errors.New("animal not found"))

		err := service.UpdateAnimal(animalID, updatedAnimal)
		assert.EqualError(t, err, "animal not found")
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)


	t.Run("should return error if save fails", func(t *testing.T) {
		mockRepo.On("FindAnimalByID", animalID).Return(existingAnimal, nil)
		mockRepo.On("SaveAnimal", existingAnimal).Return(errors.New("save failed"))

		err := service.UpdateAnimal(animalID, updatedAnimal)
		assert.EqualError(t, err, "save failed")
		mockRepo.AssertExpectations(t)
	})
}

// Teste para DeleteAnimal
func TestDeleteAnimal(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)

	animalID := uuid.New()

	t.Run("should delete animal", func(t *testing.T) {
		mockRepo.On("DeleteAnimal", animalID).Return("Animal excluído com sucesso", nil)

		msg, err := service.DeleteAnimal(animalID)
		assert.NoError(t, err)
		assert.Equal(t, "Animal excluído com sucesso", msg)
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)


	t.Run("should return error if deletion fails", func(t *testing.T) {
		mockRepo.On("DeleteAnimal", animalID).Return("", errors.New("delete failed"))

		msg, err := service.DeleteAnimal(animalID)
		assert.EqualError(t, err, "delete failed")
		assert.Empty(t, msg)
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)


	t.Run("should return error if animal not found", func(t *testing.T) {
		mockRepo.On("DeleteAnimal", animalID).Return("Animal not found", errors.New("animal not found"))

		msg, err := service.DeleteAnimal(animalID)
		assert.EqualError(t, err, "animal not found")
		assert.Equal(t, "Animal not found", msg)
		mockRepo.AssertExpectations(t)
	})
}

// Teste para GetAllAnimals
func TestGetAllAnimals(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)

	animals := []model.Animal{
		{Name: "Rex", Species: "Dog"},
		{Name: "Max", Species: "Cat"},
	}

	t.Run("should return all animals", func(t *testing.T) {
		mockRepo.On("FindAllAnimals").Return(animals, nil)

		result, err := service.GetAllAnimals()
		assert.NoError(t, err)
		assert.Equal(t, animals, result)
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)


	t.Run("should return empty list when no animals found", func(t *testing.T) {
		mockRepo.On("FindAllAnimals").Return([]model.Animal{}, nil)

		result, err := service.GetAllAnimals()
		assert.NoError(t, err)
		assert.Empty(t, result)
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)


	t.Run("should return error when fetching animals fails", func(t *testing.T) {
		mockRepo.On("FindAllAnimals").Return(nil, errors.New("fetch failed"))

		result, err := service.GetAllAnimals()
		assert.Nil(t, result)
		assert.EqualError(t, err, "fetch failed")
		mockRepo.AssertExpectations(t)
	})
}

// Teste para ValidateAnimal
func TestValidateAnimal(t *testing.T) {
	tests := []struct {
		name     string
		animal   model.Animal
		expected error
	}{
		{
			name: "valid animal",
			animal: model.Animal{
				Name:    "Rex",
				Species: "Dog",
			},
			expected: nil,
		},
		{
			name: "missing name",
			animal: model.Animal{
				Species: "Dog",
			},
			expected: errors.New("animal needs a name"),
		},
		{
			name: "missing species",
			animal: model.Animal{
				Name: "Rex",
			},
			expected: errors.New("animal needs a species"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.ValidateAnimal(tt.animal)
			assert.Equal(t, tt.expected, err)
		})
	}
}

// Teste para ValidateAnimalExists
func TestValidateAnimalExists(t *testing.T) {
	mockRepo := new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)

	animal := model.Animal{
		Name:    "Rex",
		Species: "Dog",
	}

	t.Run("should return error if animal exists", func(t *testing.T) {
		mockRepo.On("FindByUniqueAttributes", animal).Return(&animal, nil)

		err := service.ValidateAnimalExists(animal)
		assert.EqualError(t, err, "animal already exists")
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)


	t.Run("should return no error if animal does not exist", func(t *testing.T) {
		mockRepo.On("FindByUniqueAttributes", animal).Return(nil, nil)

		err := service.ValidateAnimalExists(animal)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	// Recriar o mock para evitar que o comportamento anterior afete este teste
	mockRepo = new(repository.MockAnimalRepository)
	service.SetAnimalRepository(mockRepo)


	t.Run("should return error if repository fails", func(t *testing.T) {
		mockRepo.On("FindByUniqueAttributes", animal).Return(nil, errors.New("repository error"))

		err := service.ValidateAnimalExists(animal)
		assert.EqualError(t, err, "repository error")
		mockRepo.AssertExpectations(t)
	})
}
