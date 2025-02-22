package repository

import (
	"handlersPetstore/internal/domain"
)

type petRepository struct{}

func NewPetRepository() domain.PetRepository {
	return &petRepository{}
}

func (r *petRepository) Create(pet *domain.Pet) error {
	return nil
}

func (r *petRepository) Update(pet *domain.Pet) error {
	return nil
}

func (r *petRepository) Delete(id int64) error {
	return nil
}

func (r *petRepository) GetByID(id int64) (*domain.Pet, error) {
	return nil, nil
}

func (r *petRepository) FindByStatus(status []string) ([]domain.Pet, error) {
	return nil, nil
}

func (r *petRepository) FindByTags(tags []string) ([]domain.Pet, error) {
	return nil, nil
}
