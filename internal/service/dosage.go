package service

import (
    "context"
    "errors"
    "vetblock/internal/db/model"
    "vetblock/internal/db/repository"

    "github.com/google/uuid"
)

type DosageService struct {
    repo repository.DosageRepository
}

// Cria uma nova instância do DosageService com um repositório de dosagem
func NewDosageService(repo repository.DosageRepository) *DosageService {
    return &DosageService{repo: repo}
}

func (s *DosageService) AddDosage(ctx context.Context, dosage *model.Dosage) error {
    if dosage == nil {
        return errors.New("dosagem não pode ser nula")
    }

    // Verifica se o animal existe
    animal, err := GetAnimalByID(dosage.AnimalID)
    if err != nil {
        return err
    }

    if animal == nil {
        return errors.New("animal não encontrado")
    }

    // Verifica se o medicamento existe
    medication, err := GetMedicationByID(dosage.MedicationID)
    if err != nil {
        return err
    }

    if medication == nil {
        return errors.New("medicamento não encontrado")
    }

    // Verifica se a consulta existe, se ConsultationsID não for nil
    if dosage.ConsultationID != nil {
        consultation, err := GetConsultationByID(*dosage.ConsultationID) // Desreferencia o ponteiro
        if err != nil {
            return err
        }

        if consultation == nil {
            return errors.New("consulta não encontrada")
        }
    }
	hospitazationRepository := repository.NewHospitalizationRepository()
    // Verifica se a hospitalização existe, se HospitalizationID não for nil
    if dosage.HospitalizationID != nil {
        hospitalization, err := hospitazationRepository.GetHospitalizationByID(*dosage.HospitalizationID) // Desreferencia o ponteiro
        if err != nil {
            return err
        }

        if hospitalization == nil {
            return errors.New("hospitalização não encontrada")
        }
    }

    return s.repo.Create(ctx, dosage)
}

func GetHospitalizationByID(uUID uuid.UUID) {
	panic("unimplemented")
}


// Atualiza uma dosagem existente
func (s *DosageService) UpdateDosage(ctx context.Context, dosage *model.Dosage) error {
    if dosage == nil {
        return errors.New("dosagem não pode ser nula")
    }

    existingDosage, err := s.repo.FindByID(ctx, dosage.ID)
    if err != nil {
        return err
    }
    if existingDosage == nil {
        return errors.New("dosagem não encontrada")
    }

    return s.repo.Update(ctx, dosage)
}

// Deleta uma dosagem pelo ID
func (s *DosageService) DeleteDosage(ctx context.Context, dosageID uuid.UUID) error {
    return s.repo.Delete(ctx, dosageID)
}

// Encontra uma dosagem pelo ID
func (s *DosageService) FindDosageByID(ctx context.Context, dosageID uuid.UUID) (*model.Dosage, error) {
    return s.repo.FindByID(ctx, dosageID)
}

// Encontra todas as dosagens associadas a um animal
func (s *DosageService) FindDosagesByAnimalID(ctx context.Context, animalID uuid.UUID) ([]model.Dosage, error) {
    return s.repo.FindByAnimalID(ctx, animalID)
}
