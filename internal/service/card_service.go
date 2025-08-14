package service

import (
	"github.com/stephenZ22/mini_dash/internal/model"
	"github.com/stephenZ22/mini_dash/internal/repository"
)

type CardService interface {
	CreateCard(card *model.Card) error
	GetCardByID(id uint) (*model.Card, error)
	UpdateCard(card *model.Card) error
	DeleteCard(id uint) error
}

type cardService struct {
	repo repository.CardRepository
}

func NewCardService(repo repository.CardRepository) CardService {
	return &cardService{repo: repo}
}

func (s *cardService) CreateCard(card *model.Card) error {
	return s.repo.CreateCard(card)
}

func (s *cardService) GetCardByID(id uint) (*model.Card, error) {
	return s.repo.GetCardByID(id)
}

func (s *cardService) UpdateCard(card *model.Card) error {
	return s.repo.UpdateCard(card)
}

func (s *cardService) DeleteCard(id uint) error {
	return s.repo.DeleteCard(id)
}
