package users

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Password  string         `json:"password" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (u User) String() string {
	return fmt.Sprintf("User(ID: %d, Name: %s, Email: %s, CreatedAt: %s, UpdatedAt: %s)",
		u.ID, u.Username, u.Email, u.CreatedAt.Format(time.RFC3339), u.UpdatedAt.Format(time.RFC3339))
}
