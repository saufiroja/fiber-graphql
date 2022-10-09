package service

import (
	"fiber/fiber-graphql/config/database"
	"fiber/fiber-graphql/domain"
	"fiber/fiber-graphql/entity"
)

type CustomerService struct {
	conf               database.Config
	CustomerRepository domain.CustomerRepository
}

func NewCustomerService(customerRepository domain.CustomerRepository, conf database.Config) domain.CustomerService {
	return &CustomerService{
		conf:               conf,
		CustomerRepository: customerRepository,
	}
}

func (s *CustomerService) GetAll() ([]entity.Customer, error) {
	return s.CustomerRepository.GetAll()
}

func (s *CustomerService) GetById(id uint) (entity.Customer, error) {
	return s.CustomerRepository.GetById(id)
}

func (s *CustomerService) CreateCustomer(customer *entity.Customer) (entity.Customer, error) {
	return s.CustomerRepository.CreateCustomer(customer)
}
