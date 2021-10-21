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
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
