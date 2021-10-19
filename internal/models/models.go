package models

import "time"

// Customer is the customer model
type Customer struct {
	ID        string
	FirstName string
	LastName  string
	BirthDate time.Time
	Email     string
	Gender    string
	CreateAt  time.Time
	UpdateAt  time.Time
}
