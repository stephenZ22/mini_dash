package projects

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
