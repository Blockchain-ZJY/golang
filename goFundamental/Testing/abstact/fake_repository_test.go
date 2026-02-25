// service_test.go
package order

import "testing"

func TestCreateOrder(t *testing.T) {
	fakeRepo := &FakeOrderRepository{}
	service := NewOrderService(fakeRepo)

	order := &Order{Price: 100}

	err := service.CreateOrder(order)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fakeRepo.SavedOrder != order {
		t.Fatalf("order not saved")
	}
}

func TestCreateOrder_InvalidPrice(t *testing.T) {
	fakeRepo := &FakeOrderRepository{}
	service := NewOrderService(fakeRepo)

	order := &Order{Price: -1}

	err := service.CreateOrder(order)
	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}
