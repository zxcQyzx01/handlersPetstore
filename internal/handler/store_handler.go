package handler

import (
	"encoding/json"
	"handlersPetstore/internal/domain"
	"handlersPetstore/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type StoreHandler struct {
	service *service.StoreService
}

func NewStoreHandler(service *service.StoreService) *StoreHandler {
	return &StoreHandler{service: service}
}

// @Summary Get pet inventories by status
// @Description Returns a map of status codes to quantities
// @Tags store
// @Produce json
// @Success 200 {object} map[string]int32
// @Router /store/inventory [get]
func (h *StoreHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	inventory, err := h.service.GetInventory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}

// @Summary Place an order for a pet
// @Description Place a new order in the store
// @Tags store
// @Accept json
// @Produce json
// @Param order body domain.Order true "Order object that needs to be added"
// @Success 200 {object} domain.Order
// @Router /store/order [post]
func (h *StoreHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.PlaceOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// @Summary Find purchase order by ID
// @Description For valid response try integer IDs with value >= 1
// @Tags store
// @Produce json
// @Param orderId path int true "ID of order that needs to be fetched"
// @Success 200 {object} domain.Order
// @Router /store/order/{orderId} [get]
func (h *StoreHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.ParseInt(chi.URLParam(r, "orderId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrderByID(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// @Summary Delete purchase order by ID
// @Description For valid response try integer IDs with positive integer value
// @Tags store
// @Produce json
// @Param orderId path int true "ID of the order that needs to be deleted"
// @Success 200 {string} string "Order deleted successfully"
// @Router /store/order/{orderId} [delete]
func (h *StoreHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.ParseInt(chi.URLParam(r, "orderId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteOrder(orderID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
