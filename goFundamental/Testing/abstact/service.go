// service.go
package order

import "errors"

type Order struct {
	ID    int
	Price int
}

type OrderRepository interface {
	Save(order *Order) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *Order) error {
	if order.Price < 0 {
		return errors.New("价格错误")
	}
	return s.repo.Save(order)
}
