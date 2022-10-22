package domain

import (
	"fiber/fiber-graphql/entity"

	"github.com/graphql-go/graphql"
)

type CustomerRepository interface {
	GetAll() ([]entity.Customer, error)
	GetById(id uint) (entity.Customer, error)
	CreateCustomer(customer *entity.Customer) (entity.Customer, error)
}

type CustomerService interface {
	GetAll() ([]entity.Customer, error)
	GetById(id uint) (entity.Customer, error)
	CreateCustomer(customer *entity.Customer) (entity.Customer, error)
}

type CustomerResolver interface {
	GetCustomer(params graphql.ResolveParams) (interface{}, error)
	GetCustomers(params graphql.ResolveParams) (interface{}, error)
	CreateCustomer(params graphql.ResolveParams) (interface{}, error)
}
