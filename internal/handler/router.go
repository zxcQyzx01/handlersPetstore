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
	})

	return r
}
