package repository

import "github.com/ekateryna-tln/wallester_task/internal/models"

type DatabaseRepo interface {
	GetAllCustomers() bool
	InsertCustomer(customer models.Customer) (string, error)
}
