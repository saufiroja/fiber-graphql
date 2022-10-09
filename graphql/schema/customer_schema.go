package schema

import "fiber/fiber-graphql/domain"

type customerSchema struct {
	customerResolver domain.CustomerResolver
}

func NewCustomerSchema(customerResolver domain.CustomerResolver) customerSchema {
	return customerSchema{
		customerResolver: customerResolver,
	}
}
