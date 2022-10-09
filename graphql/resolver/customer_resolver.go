package resolver

import (
	"fiber/fiber-graphql/domain"
	"fiber/fiber-graphql/entity"
	"fmt"

	"github.com/graphql-go/graphql"
)

type CustomerResolver struct {
	customerService domain.CustomerService
}

func NewCustomerResolver(customerService domain.CustomerService) domain.CustomerResolver {
	return &CustomerResolver{
		customerService: customerService,
	}
}

func (r *CustomerResolver) GetCustomer(params graphql.ResolveParams) (interface{}, error) {
	var (
		id int
		ok bool
	)

	if id, ok = params.Args["id"].(int); !ok {
		return nil, fmt.Errorf("id is required")
	}

	customer, err := r.customerService.GetById(uint(id))
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerResolver) GetCustomers(params graphql.ResolveParams) (interface{}, error) {
	customers, err := r.customerService.GetAll()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerResolver) CreateCustomer(params graphql.ResolveParams) (interface{}, error) {
	var (
		name string
		age  int
		ok   bool
	)

	if name, ok = params.Args["name"].(string); !ok {
		return nil, fmt.Errorf("name is required")
	}
	if age, ok = params.Args["age"].(int); !ok {
		return nil, fmt.Errorf("age is required")
	}

	customer := entity.Customer{
		Name: name,
		Age:  age,
	}

	newCustomer, err := r.customerService.CreateCustomer(&customer)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}
