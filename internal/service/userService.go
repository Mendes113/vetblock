package service

import (
    "errors"
    "vetblock/internal/db"
    "vetblock/internal/db/model"
)

// Cria um usuário genérico (Tutor ou Veterinarian)
func CreateUser(user interface{}) error {
    switch u := user.(type) {
    case *model.Tutor:
        if err := db.NewDb().Save(u).Error; err != nil {
            return err
        }
    case *model.Veterinarian:
        if err := db.NewDb().Save(u).Error; err != nil {
            return err
        }
    default:
        return errors.New("invalid user type")
    }
    return nil
}
