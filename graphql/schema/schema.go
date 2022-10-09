package schema

import "github.com/graphql-go/graphql"

var (
	Customer = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Customer",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"age": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
)

func (c customerSchema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"GetCustomers": &graphql.Field{
				Type:        graphql.NewList(Customer),
				Description: "Get All Customer",
				Resolve:     c.customerResolver.GetCustomers,
			},
			"GetCustomer": &graphql.Field{
				Type:        Customer,
				Description: "Get Customer By Id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: c.customerResolver.GetCustomer,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}

func (c customerSchema) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"CreateCustomer": &graphql.Field{
				Type:        Customer,
				Description: "Create Customer",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: c.customerResolver.CreateCustomer,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}
