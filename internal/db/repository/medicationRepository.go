package repository

import (
	"log"
	"time"
	"vetblock/internal/db"
	"vetblock/internal/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicationRepository struct {
	Db *gorm.DB
}

func NewMedicationRepository() *MedicationRepository {
	return &MedicationRepository{
		Db: db.NewDb(),
	}
}

// Salva um medicamento no banco de dados e retorna um erro se ocorrer
func (r *MedicationRepository) SaveMedication(medication *model.Medication) error {
	if err := r.Db.Create(medication).Error; err != nil {
		log.Print("Error saving medication:", err)
		return err
	}
	log.Print(medication)
	log.Print("Repository Saving Medication")
	return nil
}

func (r *MedicationRepository) FindByUniqueAttributes(medication *model.Medication) *model.Medication {
	var med model.Medication
	if err := r.Db.Where("name = ? AND concentration = ? AND presentation = ?", medication.Name, medication.Concentration, medication.Presentation).First(&med).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil
	}
	return &med
}

// by name
func (r *MedicationRepository) FindMedicationByName(name string) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("name = ?", name).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}
	return &medication, nil
}

func (r *MedicationRepository) FindMedicationByID(id uuid.UUID) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("id = ? AND deleted_at IS NULL", id).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}
	return &medication, nil
}

// delete medication
func (r *MedicationRepository) DeleteMedication(id string) (string, error) {
	var medication model.Medication
	if err := r.Db.Where("id = ?", id).First(&medication).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Print("Medication not found:", err)
			return "Medication not found", err
		}
		log.Print("Error finding medication:", err)
		return "Error finding medication", err
	}

	// Soft delete
	if err := r.Db.Delete(&medication).Error; err != nil {
		log.Print("Error deleting medication:", err)
		return "Error deleting medication", err
	}

	log.Print("Medication excluído com sucesso")
	return "Medication deleted successfully", nil
}

func (r *MedicationRepository) UpdateMedication(medication *model.Medication) error {
	if err := r.Db.Model(&medication).Updates(medication).Error; err != nil {
		log.Print("Error updating medication:", err)
		return err
	}
	log.Print("Medication updated successfully")
	return nil
}

// increase medication quantity
func (r *MedicationRepository) IncreaseMedicationQuantity(id uuid.UUID, quantity int) error {
	err := r.Db.Model(&model.Medication{}).Where("id = ?", id).UpdateColumn("quantity", gorm.Expr("quantity + ?", quantity)).Error
	if err != nil {
		log.Print("Error updating medication quantity:", err)
		return err
	}
	log.Print("Medication quantity updated successfully")
	return nil
}

func (r *MedicationRepository) FindAllMedications() ([]model.Medication, error) {
	var medications []model.Medication
	if err := r.Db.Find(&medications).Error; err != nil {
		log.Print("Error finding medications:", err)
		return nil, err
	}
	return medications, nil
}

// medication with the closest expiration date
func (r *MedicationRepository) FindMedicationClosestExpirationDate() (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("deleted_at IS NULL").Order("expiration_date asc").First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}
	return &medication, nil
}

// find medication that will and expire in a range of days
func (r *MedicationRepository) FindMedicationWillExpireInDays(days int) ([]model.Medication, error) {
	var medications []model.Medication
	if err := r.Db.Where("expiration_date BETWEEN ? AND ? AND deleted_at IS NULL", time.Now(), time.Now().AddDate(0, 0, days)).
		Order("expiration_date asc").Find(&medications).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}
	return medications, nil
}

// medication expired
func (r *MedicationRepository) FindMedicationExpired() (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("expiration < ? AND deleted_at IS NULL", time.Now()).
		Order("expiration desc"). // Ordena do mais recente ao mais antigo
		First(&medication).Error; err != nil {
		log.Print("Error finding expired medication:", err)
		return nil, err
	}
	return &medication, nil
}

//medications by batch number

func (r *MedicationRepository) FindMedicationByBatchNumber(batchNumber string) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("batch_number = ?", batchNumber).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}

	return &medication, nil
}

// medications by name

// medications by concentration
func (r *MedicationRepository) FindMedicationByConcentration(concentration string) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("concentration = ?", concentration).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}

	return &medication, nil
}

// medications by presentation
func (r *MedicationRepository) FindMedicationByPresentation(presentation string) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("presentation = ?", presentation).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}

	return &medication, nil
}

// medications by expiration date
func (r *MedicationRepository) FindMedicationByExpirationDate(expirationDate string) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("expiration_date = ?", expirationDate).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}

	return &medication, nil
}

// medications by price
func (r *MedicationRepository) FindMedicationByPrice(price float64) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("price = ?", price).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}

	return &medication, nil
}

// medications by quantity
func (r *MedicationRepository) FindMedicationByQuantity(quantity int) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("quantity = ?", quantity).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}

	return &medication, nil
}

// medication by manufacturer
func (r *MedicationRepository) FindMedicationByManufacturer(manufacturer string) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("manufacturer = ?", manufacturer).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}

	return &medication, nil
}

// medication by manufacturer and name
func (r *MedicationRepository) FindMedicationByManufacturerAndName(manufacturer string, name string) (*model.Medication, error) {
	var medication model.Medication
	if err := r.Db.Where("manufacturer = ? AND name = ?", manufacturer, name).First(&medication).Error; err != nil {
		log.Print("Error finding medication:", err)
		return nil, err
	}

	return &medication, nil
}

// medication by active substance
func (r *MedicationRepository) FindMedicationByActiveSubstance(activeSubstance string) ([]model.Medication, error) {
	var medications []model.Medication
	if err := r.Db.Where("JSON_CONTAINS(active_substance, ?, '$')", activeSubstance).Find(&medications).Error; err != nil {
		log.Print("Error finding medications:", err)
		return nil, err
	}

	return medications, nil
}

