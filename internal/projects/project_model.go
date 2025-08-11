package projects

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;unique"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime" default:"CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime" default:"CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
