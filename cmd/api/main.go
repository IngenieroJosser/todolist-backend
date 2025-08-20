package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todolist-backend/handlers"
	"todolist-backend/models"
	"todolist-backend/repository"
	"todolist-backend/routes"
	"todolist-backend/middlewares"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Configuración de DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la DB: ", err)
	}

	// Migraciones
	db.AutoMigrate(&models.User{})

	// Repositories
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	projectRepo := repository.NewProjectRepository(db)

	// Handlers
	userHandler := handlers.NewUserHandler(userRepo)
	taskHandler := handlers.NewTaskHandler(taskRepo)
	authHandler := handlers.NewAuthHandler(userRepo)
	projectHandler := handlers.NewProjectHandler(projectRepo)

	// Configura todas las rutas
	router := routes.SetupRoutes(userHandler, taskHandler, authHandler, projectHandler)

	// Wrap con el middleware CORS para permitir peticiones
	handlerWithCORS := middlewares.CORSMiddleware(router)

	// Servidor
	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}
