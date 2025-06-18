package client

import "github.com/xilepeng/100-Go-Mistakes/02-code-project-organization/6-interface-producer/store"

// import "github.com/xilepeng/100-go-mistakes/02-code-project-organization/6-interface-producer/store"

type customersGetter interface {
	GetAllCustomers() ([]store.Customer, error)
}
