package projects

import (
	"github.com/stephenZ22/mini_dash/pkg/logger"
	"gorm.io/gorm"
)

// TODO use unexport struct name
// TODO use exported interface name
// ProjectRepository is the interface for project repository operations

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository) Create(project *Project) error {

	logger.MiniLogger().Debug("Creating project", "project", project)
	if err := r.db.Create(project).Error; err != nil {
		return err
	}
	return nil
}
