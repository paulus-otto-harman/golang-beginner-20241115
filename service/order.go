package service

import (
	"20241115/model"
	"20241115/repository"
)

type OrderService struct {
	OrderRepo repository.Order
}

func InitOrderService(repo repository.Order) *OrderService {
	return &OrderService{OrderRepo: repo}
}

func (orderService OrderService) Create(order *model.Order) error {
	return orderService.OrderRepo.Create(order)
}
