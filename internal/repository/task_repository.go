package repository

import (
	"github.com/rina-makita0320/moving-checklist/internal/database"
	"github.com/rina-makita0320/moving-checklist/internal/model"
)

func GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	result := database.DB.Preload("Category").Find(&tasks)
	return tasks, result.Error
}

func CreateTask(task *model.Task) error {
	result := database.DB.Create(task)
	return result.Error
}

func CompleteTask(id uint) error {
	result := database.DB.Model(&model.Task{}).Where("id = ?", id).Update("completed", true)
	return result.Error
}