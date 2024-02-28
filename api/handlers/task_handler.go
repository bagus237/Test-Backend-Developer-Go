package handlers

import (
	"net/http"
	"strconv"

	"test_backend_developer_go/models"
	"test_backend_developer_go/repository"

	"github.com/gin-gonic/gin"
)

// TaskHandler adalah struct untuk menangani permintaan terkait tugas
type TaskHandler struct {
	TaskRepository *repository.TaskRepository
}

// NewTaskHandler membuat instance TaskHandler baru
func NewTaskHandler(taskRepo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{
		TaskRepository: taskRepo,
	}
}

// CreateTask membuat tugas baru
func (th *TaskHandler) CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.TaskRepository.CreateTask(&newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat tugas baru"})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// GetTasks mengambil semua tugas
func (th *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := th.TaskRepository.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil tugas"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTask mengambil detail tugas berdasarkan ID
func (th *TaskHandler) GetTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tugas tidak valid"})
		return
	}

	// Panggil repository untuk mendapatkan detail tugas
	taskDetail, err := th.TaskRepository.GetTaskByID(uint(taskID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tugas tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, taskDetail)
}

// UpdateTask memperbarui informasi tugas berdasarkan ID
func (th *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tugas tidak valid"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask.ID = int(taskID) // set ID tugas dengan ID yang diterima

	if err := th.TaskRepository.UpdateTask(&updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui tugas"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask menghapus tugas berdasarkan ID
func (th *TaskHandler) DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tugas tidak valid"})
		return
	}

	task := &models.Task{ID: int(taskID)}

	if err := th.TaskRepository.DeleteTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus tugas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tugas berhasil dihapus"})
}
