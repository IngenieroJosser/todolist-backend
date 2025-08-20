package repository

import (
	"todolist-backend/models"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) CreateProject(project *models.Project) error {
	return r.DB.Create(project).Error
}

func (r *ProjectRepository) FindAllProject() ([]models.Project, error) {
	var searchingProjects []models.Project
	result := r.DB.Find(&searchingProjects)
	return searchingProjects, result.Error
}

func (r *ProjectRepository) FindProjectByID(projectId uint) (models.Project, error) {
	var projectFound models.Project
	result := r.DB.First(&projectFound, projectId)
	return projectFound, result.Error
}

func (r *ProjectRepository) UpdateProject(projectData *models.Project) error {
	return r.DB.Save(projectData).Error
}

func (r *ProjectRepository) DeleteProject(projectId uint) error {
	return r.DB.Delete(&models.Project{}, projectId).Error
}