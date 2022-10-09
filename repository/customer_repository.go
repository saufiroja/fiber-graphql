package repository

import (
	"fiber/fiber-graphql/domain"
	"fiber/fiber-graphql/entity"

	"gorm.io/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{
		DB: db,
	}
}

func (r *customerRepository) GetAll() ([]entity.Customer, error) {
	var customers []entity.Customer
	_ = r.DB.Find(&customers).Error
	return customers, nil
}

func (r *customerRepository) GetById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	err := r.DB.First(&customer, id).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *customerRepository) CreateCustomer(customer *entity.Customer) (entity.Customer, error) {
	err := r.DB.Create(customer).Error
	if err != nil {
		return *customer, err
	}
	return *customer, nil
}
