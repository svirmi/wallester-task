package dbrepo

import (
	"context"
	"github.com/ekateryna-tln/wallester_task/internal/models"
	"github.com/gofrs/uuid"
	"log"
	"time"
)

func (dbRepo *postgresDBRepo) GetAllCustomers() bool {
	return true
}

func (dbRepo *postgresDBRepo) InsertCustomer(c models.Customer) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	stmt := `insert into customers
			(id ,first_name, last_name, birthdate, email, gender, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = dbRepo.DB.ExecContext(ctx, stmt,
		uuid,
		c.FirstName,
		c.LastName,
		c.Birthdate,
		c.Email,
		c.Gender,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}
