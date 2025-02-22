package handler

import (
	"handlersPetstore/internal/middleware"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewRouter(petHandler *PetHandler, userHandler *UserHandler, storeHandler *StoreHandler) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	// Pet routes
	r.Route("/v2/pet", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Post("/", petHandler.CreatePet)
		r.Put("/", petHandler.UpdatePet)
		r.Get("/{petId}", petHandler.GetPetByID)
		r.Delete("/{petId}", petHandler.DeletePet)
		r.Get("/findByStatus", petHandler.FindPetsByStatus)
		r.Get("/findByTags", petHandler.FindPetsByTags)
	})

	// Store routes
	r.Route("/v2/store", func(r chi.Router) {
		r.Get("/inventory", storeHandler.GetInventory)
		r.Post("/order", storeHandler.PlaceOrder)
		r.Get("/order/{orderId}", storeHandler.GetOrderByID)
		r.Delete("/order/{orderId}", storeHandler.DeleteOrder)
	})

	// User routes
	r.Route("/v2/user", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/createWithArray", userHandler.CreateUsersWithArrayInput)
		r.Get("/login", userHandler.Login)
		r.Get("/logout", userHandler.Logout)
		r.Get("/{username}", userHandler.GetUserByName)
		r.Put("/{username}", userHandler.UpdateUser)
		r.Delete("/{username}", userHandler.DeleteUser)
	})

	return r
}
