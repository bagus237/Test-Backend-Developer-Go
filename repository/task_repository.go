package repository

import (
	"test_backend_developer_go/models"

	"github.com/jinzhu/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

// NewTaskRepository inisialisasi repository tugas baru
func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

// CreateTask membuat tugas baru dalam database
func (tr *TaskRepository) CreateTask(task *models.Task) error {
	return tr.DB.Create(task).Error
}

// GetAllTasks mengambil semua tugas dari database
func (tr *TaskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := tr.DB.Find(&tasks).Error
	return tasks, err
}

// GetTaskByID mengambil tugas berdasarkan ID dari database
func (tr *TaskRepository) GetTaskByID(taskID uint) (*models.Task, error) {
	var task models.Task
	err := tr.DB.Where("id = ?", taskID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// UpdateTask mengupdate tugas dalam database
func (tr *TaskRepository) UpdateTask(task *models.Task) error {
	return tr.DB.Save(task).Error
}

// DeleteTask menghapus tugas dari database
func (tr *TaskRepository) DeleteTask(task *models.Task) error {
	return tr.DB.Delete(task).Error
}
