package service

import (
	"handlersPetstore/internal/domain"
)

type StoreService struct {
	repo domain.StoreRepository
}

func NewStoreService(repo domain.StoreRepository) *StoreService {
	return &StoreService{repo: repo}
}

func (s *StoreService) GetInventory() (map[string]int32, error) {
	return s.repo.GetInventory()
}

func (s *StoreService) PlaceOrder(order *domain.Order) error {
	return s.repo.PlaceOrder(order)
}

func (s *StoreService) GetOrderByID(orderID int64) (*domain.Order, error) {
	return s.repo.GetOrderByID(orderID)
}

func (s *StoreService) DeleteOrder(orderID int64) error {
	return s.repo.DeleteOrder(orderID)
}
