package cards

import (
	"fmt"
	"time"

	"github.com/stephenZ22/mini_dash/internal/projects"
	"github.com/stephenZ22/mini_dash/internal/users"
	"gorm.io/gorm"
)

// id SERIAL PRIMARY KEY,
// name VARCHAR(255),
// description  TEXT,
// card_type SMAllINT,
// parent_id INT REFERENCES cards(id) ON DELETE SET NULL,
// creater_id INT REFERENCES users(id)

type CardType int

const (
	Feature CardType = iota + 1
	Task
	Bug
)

type Card struct {
	ID          uint              `json:"id" gorm:"primaryKey"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	CardType    uint              `json:"card_type"`
	CreaterID   uint              `json:"creater_id"`
	Creater     *users.User       `gorm:"foreignKey:CreaterID;references:ID"`
	ProjectID   *uint             `json:"project_id"`
	Project     *projects.Project `gorm:"foreignKey:ProjectID;references:ID"`
	ParentID    *uint             `json:"parent_id"`
	Parent      *Card             `gorm:"foreignKey:ParentID;references:ID"`
	Children    []Card            `gorm:"foreignKey:ParentID;references:ID"`
	CreatedAt   time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt    `json:"deleted_at,omitempty" gorm:"index"`
}

func (c Card) String() string {
	return fmt.Sprintf("Card(ID: %d, Name: %s, Type: %d, CreatedAt: %s, UpdatedAt: %s)",
		c.ID, c.Name, c.CardType, c.CreatedAt.Format(time.RFC3339), c.UpdatedAt.Format(time.RFC3339))
}
