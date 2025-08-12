package projects

import "gorm.io/gorm"

type ProjectService struct {
	repo *ProjectRepository
}

func NewProjectService(repo *ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) CreateProject(name, desc string) error {
	project := &Project{
		Name:        name,
		Description: desc,
	}

	if err := s.repo.Create(project); err != nil {
		return err
	}
	return nil
}

func (s *ProjectService) GetProject(id uint) (*Project, error) {
	project, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *ProjectService) UpdateProject(id uint, project *Project) error {
	if err := s.repo.Update(id, project); err != nil {
		return err
	}
	return nil
}

func (s *ProjectService) DeleteProject(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
func (s *ProjectService) ListProjects(pageSize, pageNum int) ([]Project, error) {
	if pageSize < 0 || pageNum < 0 {
		return nil, gorm.ErrInvalidValue
	}

	if pageNum == 0 {
		pageNum = 1
	}

	if pageSize == 0 {
		pageSize = 10 // Default page size
	}

	projects, err := s.repo.List(pageSize, pageNum)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
