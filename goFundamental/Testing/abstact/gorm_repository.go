// gorm_repository.go
package order

import "gorm.io/gorm"

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order *Order) error {
	return r.db.Create(order).Error
}
