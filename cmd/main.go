package main

import (
	"log"
	"os"

	"github.com/ahmadzakyarifin/gin-jwt-auth/config"
	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/handler"
	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/repository"
	"github.com/ahmadzakyarifin/gin-jwt-auth/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment system")
	}

	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}
	defer db.Close()

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewUserHandler(authService)

	server := gin.Default()

	authRoutes := server.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("http://localhost:", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
