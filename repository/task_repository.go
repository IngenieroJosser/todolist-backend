package repository

import (
	"todolist-backend/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) FindAllTask() ([]models.Task, error) {
	var searchingTasks []models.Task
	result := r.DB.Find(&searchingTasks)
	return searchingTasks, result.Error
}

func (r *TaskRepository) FindTaskByID(taskId uint) (models.Task, error) {
	var taskFound models.Task
	result := r.DB.First(&taskFound, taskId)
	return taskFound, result.Error
}

func (r *TaskRepository) UpdateTask(taskData *models.Task) error {
	return r.DB.Save(taskData).Error
}

func (r *TaskRepository) DeleteTask(taskId uint) error {
	return r.DB.Delete(&models.Task{}, taskId).Error
}