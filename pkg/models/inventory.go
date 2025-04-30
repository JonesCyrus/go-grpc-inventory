package models

import "github.com/google/uuid"

type Inventory struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name,omitempty"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
}
