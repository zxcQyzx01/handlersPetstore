package main

import (
	_ "handlersPetstore/docs" // Import swagger docs
	"handlersPetstore/internal/handler"
	"handlersPetstore/internal/repository"
	"handlersPetstore/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Petstore API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2
func main() {
	// Инициализация репозиториев
	petRepo := repository.NewPetRepository()
	userRepo := repository.NewUserRepository()
	storeRepo := repository.NewStoreRepository()

	// Инициализация сервисов
	petService := service.NewPetService(petRepo)
	userService := service.NewUserService(userRepo)
	storeService := service.NewStoreService(storeRepo)

	// Инициализация хендлеров
	petHandler := handler.NewPetHandler(petService)
	userHandler := handler.NewUserHandler(userService)
	storeHandler := handler.NewStoreHandler(storeService)

	// Инициализация роутера
	router := handler.NewRouter(petHandler, userHandler, storeHandler)

	// Запуск сервера
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func NewRouter(petHandler *handler.PetHandler, userHandler *handler.UserHandler, storeHandler *handler.StoreHandler) *chi.Mux {
	r := chi.NewRouter()

	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return r
}
