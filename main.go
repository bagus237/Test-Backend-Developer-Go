package main

import (
	"test_backend_developer_go/api/handlers"
	"test_backend_developer_go/config"
	"test_backend_developer_go/migrations"
	"test_backend_developer_go/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi koneksi database
	db, err := config.NewDBConnection()
	if err != nil {
		panic("Gagal menghubungkan database")
	}
	defer db.Close()

	// Auto migrate struktur database
	err = migrations.AutoMigrate(db)
	if err != nil {
		panic("Gagal melakukan migrasi database")
	}

	// Inisialisasi repository
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// Inisialisasi handler dengan menyediakan repository
	userHandler := handlers.NewUserHandler(userRepo)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	// Konfigurasi router
	router := gin.Default()

	// Route untuk pengguna
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/", userHandler.GetUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Route untuk tugas
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.POST("/", taskHandler.CreateTask)
		taskRoutes.GET("/", taskHandler.GetTasks)
		taskRoutes.GET("/:id", taskHandler.GetTask)
		taskRoutes.PUT("/:id", taskHandler.UpdateTask)
		taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
	}

	// Jalankan server HTTP
	router.Run(":8080")
}
