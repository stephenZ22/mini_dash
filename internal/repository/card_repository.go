package repository

import (
	"github.com/stephenZ22/mini_dash/internal/model"
	"gorm.io/gorm"
)

type CardRepository interface {
	CreateCard(card *model.Card) error
	GetCardByID(id uint) (*model.Card, error)
	UpdateCard(card *model.Card) error
	DeleteCard(id uint) error
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) CreateCard(card *model.Card) error {
	return r.db.Create(card).Error
}

func (r *cardRepository) GetCardByID(id uint) (*model.Card, error) {
	var card model.Card
	if err := r.db.First(&card, id).Error; err != nil {
		return nil, err
	}
	return &card, nil
}

func (r *cardRepository) UpdateCard(card *model.Card) error {
	return r.db.Save(card).Error
}

func (r *cardRepository) DeleteCard(id uint) error {
	return r.db.Delete(&model.Card{}, id).Error
}
