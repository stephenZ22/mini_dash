package cards

type CardService interface {
	CreateCard(card *Card) error
	GetCardByID(id uint) (*Card, error)
	UpdateCard(card *Card) error
	DeleteCard(id uint) error
}

type cardService struct {
	repo CardRepository
}

func NewCardService(repo CardRepository) CardService {
	return &cardService{repo: repo}
}

func (s *cardService) CreateCard(card *Card) error {
	return s.repo.CreateCard(card)
}

func (s *cardService) GetCardByID(id uint) (*Card, error) {
	return s.repo.GetCardByID(id)
}

func (s *cardService) UpdateCard(card *Card) error {
	return s.repo.UpdateCard(card)
}

func (s *cardService) DeleteCard(id uint) error {
	return s.repo.DeleteCard(id)
}
