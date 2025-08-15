package repository

import (
	"github.com/stephenZ22/mini_dash/internal/model"
	"github.com/stephenZ22/mini_dash/pkg/logger"
	"gorm.io/gorm"
)

// TODO use unexport struct name
// TODO use exported interface name
// ProjectRepository is the interface for project repository operations

type ProjectRepository interface {
	Create(project *model.Project) error
	GetByID(id uint) (*model.Project, error)
	Update(id uint, project *model.Project) error
	Delete(id uint) error
	List(page_size, page_num int) ([]model.Project, error)
}
type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) Create(project *model.Project) error {

	logger.MiniLogger().Info("project Creating", "project:", project)
	if err := r.db.Create(project).Error; err != nil {
		return err
	}
	return nil
}

func (r *projectRepository) GetByID(id uint) (*model.Project, error) {
	var project model.Project
	if err := r.db.First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) Update(id uint, project *model.Project) error {
	logger.MiniLogger().Info("project Updating", "project:", project)
	// 找到对应的project并更新
	if err := r.db.First(&model.Project{}, id).Error; err != nil {
		return err
	}

	if err := r.db.Model(&model.Project{}).Where("id = ?", id).Updates(project).Error; err != nil {
		return err
	}
	logger.MiniLogger().Info("project Updated", "project:", project)
	return nil
}

func (r *projectRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Project{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *projectRepository) List(page_size, page_num int) ([]model.Project, error) {
	var projects []model.Project

	if page_size < 0 || page_num < 0 {
		return nil, gorm.ErrInvalidValue
	}

	if page_num == 0 {
		page_num = 1
	}

	if page_size == 0 {
		page_size = 20
	}
	if err := r.db.Offset((page_num - 1) * page_size).Limit(page_size).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
