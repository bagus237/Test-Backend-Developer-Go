package handlers

import (
	"net/http"
	"strconv"

	"test_backend_developer_go/models"
	"test_backend_developer_go/repository"

	"github.com/gin-gonic/gin"
)

// TaskHandler adalah struct untuk menangani permintaan terkait task
type TaskHandler struct {
	TaskRepository *repository.TaskRepository
}

// NewTaskHandler membuat instance TaskHandler baru
func NewTaskHandler(taskRepo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{
		TaskRepository: taskRepo,
	}
}

// CreateTask membuat task baru
func (th *TaskHandler) CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.TaskRepository.CreateTask(&newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat task baru"})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// GetTasks mengambil semua task
func (th *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := th.TaskRepository.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil task"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTask mengambil detail task berdasarkan ID
func (th *TaskHandler) GetTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID task tidak valid"})
		return
	}

	// Panggil repository untuk mendapatkan detail task
	taskDetail, err := th.TaskRepository.GetTaskByID(uint(taskID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, taskDetail)
}

// UpdateTask memperbarui informasi task berdasarkan ID
func (th *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID task tidak valid"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask.ID = int(taskID) // set ID task dengan ID yang diterima

	if err := th.TaskRepository.UpdateTask(&updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui task"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask menghapus task berdasarkan ID
func (th *TaskHandler) DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID task tidak valid"})
		return
	}

	task := &models.Task{ID: int(taskID)}

	if err := th.TaskRepository.DeleteTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task berhasil dihapus"})
}
