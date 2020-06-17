package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// BaseModel .
type BaseModel struct {
	gorm.Model
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
