package model

import "github.com/google/uuid"



type ImageModel struct {
    ID    uuid.UUID   `gorm:"primaryKey" type:"uuid"`
    Image []byte `gorm:"type:bytea"`
}
