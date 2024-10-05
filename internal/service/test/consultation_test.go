package service_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"
	"vetblock/internal/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do repositório
type MockConsultationRepo struct {
	mock.Mock
}

var _ repository.ConsultationRepository = (*MockConsultationRepo)(nil)

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

func (m *MockConsultationRepo) FindConsultationByVeterinaryCRVM(ctx context.Context, crvm string) ([]model.Consultation, error) {
	args := m.Called(ctx, crvm)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Consultation), args.Error(1)
}

// Adicionando outros métodos da interface caso necessário
// Para que a interface seja cumprida, todos os métodos são mockados, mesmo que não utilizados diretamente no teste.

func (m *MockConsultationRepo) FindConsultationByAnimalID(ctx context.Context, animalID uuid.UUID) ([]model.Consultation, error) {
	args := m.Called(ctx, animalID)
	return args.Get(0).([]model.Consultation), args.Error(1)
}

func (m *MockConsultationRepo) FindConsultationByDate(ctx context.Context, date string) ([]model.Consultation, error) {
	args := m.Called(ctx, date)
	return args.Get(0).([]model.Consultation), args.Error(1)
}

func (m *MockConsultationRepo) FindConsultationByDateRange(ctx context.Context, startDate, endDate string) ([]model.Consultation, error) {
	args := m.Called(ctx, startDate, endDate)
	return args.Get(0).([]model.Consultation), args.Error(1)
}

func (m *MockConsultationRepo) FindConsultationByAnimalIDAndDateRange(ctx context.Context, animalID uuid.UUID, startDate, endDate string) ([]model.Consultation, error) {
	args := m.Called(ctx, animalID, startDate, endDate)
	return args.Get(0).([]model.Consultation), args.Error(1)
}

func (m *MockConsultationRepo) FindConsultationByAnimalIDAndDate(ctx context.Context, animalID uuid.UUID, date string) ([]model.Consultation, error) {
	args := m.Called(ctx, animalID, date)
	return args.Get(0).([]model.Consultation), args.Error(1)
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

// Testes para AddConsultation
func TestAddConsultation(t *testing.T) {
	mockRepo := new(MockConsultationRepo)

	// Cenário de sucesso
	t.Run("Consulta adicionada com sucesso", func(t *testing.T) {
		consultation := &model.Consultation{
			ID:               uuid.New(),
			AnimalID:         uuid.New(),
			CRVM:             "valid-crvm",
			ConsultationDate: model.CustomDate{Time: time.Now()},
			ConsultationHour: "14:30",  // Definindo a hora da consulta
		}
	
		// Mock: Nenhuma consulta existente com esse ID
		mockRepo.On("FindConsultationByID", mock.Anything, consultation.ID).Return(nil, nil)
		
		// Mock: Retornar nenhuma consulta conflitante no mesmo dia
		mockRepo.On("FindConsultationByDate", mock.Anything, consultation.ConsultationDate.String()).Return([]model.Consultation{}, nil)
	
		// Mock: Salvar a consulta com sucesso
		mockRepo.On("SaveConsultation", mock.Anything, consultation).Return(nil)
	
		// Chamada do serviço para adicionar a consulta
		err := service.AddConsultation(mockRepo, consultation, MockGetVeterinaryByCRVM, MockGetAnimalByID)
	
		// Verificar que nenhum erro foi retornado
		assert.NoError(t, err)
	
		// Verificar que as expectativas do mock foram cumpridas
		mockRepo.AssertExpectations(t)
	})
	
	

	// Cenário onde a consulta já existe
	t.Run("Consulta já existe", func(t *testing.T) {
		consultation := &model.Consultation{ID: uuid.New()}

		mockRepo.On("FindConsultationByID", mock.Anything, consultation.ID).Return(consultation, nil)

		err := service.AddConsultation(mockRepo, consultation, MockGetVeterinaryByCRVM, MockGetAnimalByID)
		assert.EqualError(t, err, "consulta já existe")
	})

	// Veterinário não encontrado
	t.Run("Veterinário não encontrado", func(t *testing.T) {
		consultation := &model.Consultation{
			ID:       uuid.New(),
			CRVM:     "invalid-crvm",
			AnimalID: uuid.New(),
		}

		// O veterinário não será encontrado, então SaveConsultation não deve ser chamado
		mockRepo.On("FindConsultationByID", mock.Anything, consultation.ID).Return(nil, nil)

		err := service.AddConsultation(mockRepo, consultation, MockGetVeterinaryByCRVM, MockGetAnimalByID)
		assert.EqualError(t, err, "veterinário não encontrado")

		// Verifica que SaveConsultation NÃO foi chamado
		mockRepo.AssertNotCalled(t, "SaveConsultation", mock.Anything, consultation)
	})
}

func TestGetNextConsultationByVeterinaryCRVM(t *testing.T) {
	crvm := "valid-crvm"

	// Definir manualmente os UUIDs para garantir consistência
	uuid1 := uuid.MustParse("6a364a37-0618-42f7-9069-e91239f352ed") // UUID fixo para consulta 1
	uuid2 := uuid.MustParse("cc8ce86a-0009-4a90-81bb-fb02a680d3f7") // UUID fixo para consulta 2

	// Criando uma lista de consultas com datas futuras
	consultation1 := model.Consultation{
		ID:               uuid1, // UUID da consulta 1
		CRVM:             crvm,
		ConsultationDate: model.CustomDate{Time: time.Now().Add(24 * time.Hour)}, // Amanhã
	}
	consultation2 := model.Consultation{
		ID:               uuid2, // UUID da consulta 2
		CRVM:             crvm,
		ConsultationDate:  model.CustomDate{Time:time.Now().Add(48 * time.Hour)}, // Dois dias depois
	}

	t.Run("Próxima consulta encontrada com sucesso", func(t *testing.T) {
		mockRepo := new(MockConsultationRepo) // Criação do mock individualmente para o sub-teste
		// Definindo o mock para retornar as consultas
		mockRepo.On("FindConsultationByVeterinaryCRVM", mock.Anything, crvm).Return([]model.Consultation{consultation1, consultation2}, nil)

		// Chame a função
		nextConsultation, err := service.GetNextConsultationByVeterinaryCRVM(mockRepo, crvm)

		// Verifique se não houve erro
		assert.NoError(t, err)

		assert.NotNil(t, nextConsultation)
		log.Println(nextConsultation)
		assert.Equal(t, uuid1, nextConsultation.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Nenhuma consulta encontrada", func(t *testing.T) {
		mockRepo := new(MockConsultationRepo) // Criação do mock individualmente para o sub-teste
		// Nenhuma consulta encontrada
		mockRepo.On("FindConsultationByVeterinaryCRVM", mock.Anything, crvm).Return([]model.Consultation{}, nil)

		// Chame a função
		nextConsultation, err := service.GetNextConsultationByVeterinaryCRVM(mockRepo, crvm)

		// Verifique se a consulta retornada é nil
		assert.Nil(t, nextConsultation)
		// Verifique se o erro corresponde à mensagem esperada
		assert.EqualError(t, err, "nenhuma consulta encontrada")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Erro ao buscar consultas", func(t *testing.T) {
		mockRepo := new(MockConsultationRepo) // Criação do mock individualmente para o sub-teste
		// Simula um erro na busca de consultas
		mockRepo.On("FindConsultationByVeterinaryCRVM", mock.Anything, crvm).Return(nil, errors.New("erro ao buscar consultas"))

		// Chame a função
		nextConsultation, err := service.GetNextConsultationByVeterinaryCRVM(mockRepo, crvm)

		// Verifique se a consulta retornada é nil
		assert.Nil(t, nextConsultation)
		// Verifique se o erro corresponde à mensagem esperada
		assert.EqualError(t, err, "erro ao buscar consultas")
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateConsultation(t *testing.T) {
	mockRepo := new(MockConsultationRepo)
	id := uuid.MustParse("6a364a37-0618-42f7-9069-e91239f352ed")
	updatedConsultation := &model.Consultation{
		AnimalID:               uuid.New(),
		CRVM:                   "valid-crvm",
		ConsultationDate:        model.CustomDate{Time:time.Now().Add(24 * time.Hour)},
		ConsultationDescription: "Updated description",
		ConsultationType:       "Check-up",
		ConsultationPrescription: "Updated prescription",
		ConsultationPrice:      100.50,
	}

	existingConsultation := &model.Consultation{
		ID:                     id,
		AnimalID:               uuid.New(),
		CRVM:                   "valid-crvm",
		ConsultationDate:        model.CustomDate{Time:time.Now()},
		ConsultationDescription: "Old description",
		ConsultationType:       "Consultation",
		ConsultationPrescription: "Old prescription",
		ConsultationPrice:      80.00,
	}

	t.Run("Consulta atualizada com sucesso", func(t *testing.T) {
		mockRepo.On("FindConsultationByID", mock.Anything, id).Return(existingConsultation, nil)
		mockRepo.On("SaveConsultation", mock.Anything, mock.Anything).Return(nil)

		err := service.UpdateConsultation(mockRepo, id, updatedConsultation)
		assert.NoError(t, err)

		// Verifica se os campos foram atualizados
		assert.Equal(t, updatedConsultation.AnimalID, existingConsultation.AnimalID)
		assert.Equal(t, updatedConsultation.CRVM, existingConsultation.CRVM)
		assert.Equal(t, updatedConsultation.ConsultationDate, existingConsultation.ConsultationDate)
		assert.Equal(t, updatedConsultation.ConsultationDescription, existingConsultation.ConsultationDescription)
		assert.Equal(t, updatedConsultation.ConsultationType, existingConsultation.ConsultationType)
		assert.Equal(t, updatedConsultation.ConsultationPrescription, existingConsultation.ConsultationPrescription)
		assert.Equal(t, updatedConsultation.ConsultationPrice, existingConsultation.ConsultationPrice)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Erro ao encontrar consulta", func(t *testing.T) {
		mockRepo := new(MockConsultationRepo) 
		mockRepo.On("FindConsultationByID", mock.Anything, id).Return(nil, errors.New("consulta não encontrada"))

		err := service.UpdateConsultation(mockRepo, id, updatedConsultation)
		assert.EqualError(t, err, "consulta não encontrada")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Erro ao salvar consulta", func(t *testing.T) {
		mockRepo := new(MockConsultationRepo) 
		mockRepo.On("FindConsultationByID", mock.Anything, id).Return(existingConsultation, nil)
		mockRepo.On("SaveConsultation", mock.Anything, mock.Anything).Return(errors.New("erro ao salvar"))

		err := service.UpdateConsultation(mockRepo, id, updatedConsultation)
		assert.EqualError(t, err, "erro ao salvar")
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteConsultation(t *testing.T) {
	mockRepo := new(MockConsultationRepo)
	id := uuid.MustParse("6a364a37-0618-42f7-9069-e91239f352ed")

	existingConsultation := &model.Consultation{
		ID:                     id,
		AnimalID:               uuid.New(),
		CRVM:                   "valid-crvm",
		ConsultationDate:        model.CustomDate{Time:time.Now()},
		ConsultationDescription: "Some description",
		ConsultationType:       "Consultation",
		ConsultationPrescription: "Some prescription",
		ConsultationPrice:      80.00,
	}

	t.Run("Consulta excluída com sucesso", func(t *testing.T) {
		mockRepo.On("FindConsultationByID", mock.Anything, id).Return(existingConsultation, nil)
		mockRepo.On("DeleteConsultation", mock.Anything, id).Return(nil)

		err := service.DeleteConsultation(mockRepo, id)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Erro ao encontrar consulta", func(t *testing.T) {
		mockRepo := new(MockConsultationRepo) 
		mockRepo.On("FindConsultationByID", mock.Anything, id).Return(nil, errors.New("consulta não encontrada"))

		err := service.DeleteConsultation(mockRepo, id)
		assert.EqualError(t, err, "consulta não encontrada")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Erro ao excluir consulta", func(t *testing.T) {
		mockRepo := new(MockConsultationRepo) 
		mockRepo.On("FindConsultationByID", mock.Anything, id).Return(existingConsultation, nil)
		mockRepo.On("DeleteConsultation", mock.Anything, id).Return(errors.New("erro ao excluir"))

		err := service.DeleteConsultation(mockRepo, id)
		assert.EqualError(t, err, "erro ao excluir")
		mockRepo.AssertExpectations(t)
	})
}
