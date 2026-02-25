// fake_repository_test.go
package order

type FakeOrderRepository struct {
	SavedOrder *Order
	SaveErr    error
}

func (f *FakeOrderRepository) Save(order *Order) error {
	f.SavedOrder = order
	return f.SaveErr
}
