package service

import (
	"handlersPetstore/internal/domain"
)

type PetService struct {
	repo domain.PetRepository
}

func NewPetService(repo domain.PetRepository) *PetService {
	return &PetService{repo: repo}
}

func (s *PetService) CreatePet(pet *domain.Pet) error {
	return s.repo.Create(pet)
}

func (s *PetService) UpdatePet(pet *domain.Pet) error {
	return s.repo.Update(pet)
}

func (s *PetService) DeletePet(id int64) error {
	return s.repo.Delete(id)
}

func (s *PetService) GetPetByID(id int64) (*domain.Pet, error) {
	return s.repo.GetByID(id)
}

func (s *PetService) FindPetsByStatus(status []string) ([]domain.Pet, error) {
	return s.repo.FindByStatus(status)
}

func (s *PetService) FindPetsByTags(tags []string) ([]domain.Pet, error) {
	return s.repo.FindByTags(tags)
}
