package cards

import "gorm.io/gorm"

type CardRepository interface {
	CreateCard(card *Card) error
	GetCardByID(id uint) (*Card, error)
	UpdateCard(card *Card) error
	DeleteCard(id uint) error
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) CreateCard(card *Card) error {
	return r.db.Create(card).Error
}

func (r *cardRepository) GetCardByID(id uint) (*Card, error) {
	var card Card
	if err := r.db.First(&card, id).Error; err != nil {
		return nil, err
	}
	return &card, nil
}

func (r *cardRepository) UpdateCard(card *Card) error {
	return r.db.Save(card).Error
}

func (r *cardRepository) DeleteCard(id uint) error {
	return r.db.Delete(&Card{}, id).Error
}
