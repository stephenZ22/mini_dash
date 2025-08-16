package service

import (
	"github.com/stephenZ22/mini_dash/internal/model"
	"github.com/stephenZ22/mini_dash/internal/repository"
	"github.com/stephenZ22/mini_dash/pkg/logger"
	"gorm.io/gorm"
)

type ProjectService interface {
	CreateProject(name, desc string) error
	GetProject(id uint) (*model.Project, error)
	UpdateProject(id uint, project *model.Project) error
	DeleteProject(id uint) error
	ListProjects(pageSize, pageNum int) ([]model.Project, error)
}
type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{
		repo: repo,
	}
}

func (s *projectService) CreateProject(name, desc string) error {
	logger.MiniLogger().Infof("create project name: %s, description: %s", name, desc)
	project := &model.Project{
		Name:        name,
		Description: desc,
	}

	if err := s.repo.Create(project); err != nil {
		return err
	}
	return nil
}

func (s *projectService) GetProject(id uint) (*model.Project, error) {
	project, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *projectService) UpdateProject(id uint, project *model.Project) error {
	if err := s.repo.Update(id, project); err != nil {
		return err
	}
	return nil
}

func (s *projectService) DeleteProject(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
func (s *projectService) ListProjects(pageSize, pageNum int) ([]model.Project, error) {
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
