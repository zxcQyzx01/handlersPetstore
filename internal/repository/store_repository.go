package repository

import (
	"handlersPetstore/internal/domain"
)

type storeRepository struct{}

func NewStoreRepository() domain.StoreRepository {
	return &storeRepository{}
}

func (r *storeRepository) GetInventory() (map[string]int32, error) {
	return nil, nil
}

func (r *storeRepository) PlaceOrder(order *domain.Order) error {
	return nil
}

func (r *storeRepository) GetOrderByID(orderID int64) (*domain.Order, error) {
	return nil, nil
}

func (r *storeRepository) DeleteOrder(orderID int64) error {
	return nil
}
