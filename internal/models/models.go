package models

import (
	"time"
)

// Customer is the customer model
type Customer struct {
	Uuid          string
	FirstName     string
	LastName      string
	Birthdate     time.Time
	BirthdateForm string
	Email         string
	Gender        string
	SearchField   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// EqualBase compares base fields of customers structures (FirstName, LastName, Birthdate, Email, Gender),
// excluded auto generated fields like Uuid, BirthdateForm, SearchField, CreatedAt, UpdatedAt
func (c *Customer) EqualBase(customer Customer) bool {
	if c.FirstName == customer.FirstName &&
		c.LastName == customer.LastName &&
		c.Birthdate == customer.Birthdate &&
		c.Email == customer.Email &&
		c.Gender == customer.Gender {
		return true
	}
	return false
}
