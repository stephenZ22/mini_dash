package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;unique"`
	Description string         `json:"description"`
	CreaterID   uint           `json:"user_id" gorm:"column:user_id"`
	Creater     *User          `gorm:"foreignKey:CreaterID;references:ID"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime" default:"CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime" default:"CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (p Project) String() string {
	return fmt.Sprintf("Project(ID: %d, Name: %s, Description: %s, CreatedAt: %s, UpdatedAt: %s)",
		p.ID, p.Name, p.Description, p.CreatedAt.Format(time.RFC3339), p.UpdatedAt.Format(time.RFC3339))
}
