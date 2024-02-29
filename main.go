package main

import (
	"net/http"
	"test_backend_developer_go/api/handlers"
	"test_backend_developer_go/config"
	"test_backend_developer_go/migrations"
	"test_backend_developer_go/repository"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
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

	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	// Inisialisasi oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	// Konfigurasi router
	router := gin.Default()

	// Route untuk autentikasi
	authRoutes := router.Group("/oauth2")
	{
		authRoutes.GET("/token", ginserver.HandleTokenRequest)
	}
	//authMiddleware := middleware.HandleTokenVerificationError()

	// Route untuk test autentikasi
	api := router.Group("/api")
	{
		api.Use(ginserver.HandleTokenVerify())
		api.GET("/test", func(c *gin.Context) {
			ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
			if exists {
				c.JSON(http.StatusOK, ti)
				return
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				return
			}
		})
	}

	// Route untuk pengguna
	userRoutes := router.Group("/users")
	userRoutes.Use(ginserver.HandleTokenVerify())

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
		taskRoutes.Use(ginserver.HandleTokenVerify())
		taskRoutes.POST("/", taskHandler.CreateTask)
		taskRoutes.GET("/", taskHandler.GetTasks)
		taskRoutes.GET("/:id", taskHandler.GetTask)
		taskRoutes.PUT("/:id", taskHandler.UpdateTask)
		taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
	}

	// Jalankan server HTTP
	router.Run(":8080")
}
