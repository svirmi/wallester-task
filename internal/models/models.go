package models

import "time"

// Customer is the customer model
type Customer struct {
	ID        string
	FirstName string
	LastName  string
	Birthdate string
	Email     string
	Gender    string
	CreateAt  time.Time
	UpdateAt  time.Time
}
