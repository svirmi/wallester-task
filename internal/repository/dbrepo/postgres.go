package dbrepo

import (
	"context"
	"fmt"
	"github.com/ekateryna-tln/wallester_task/internal/models"
	"github.com/gofrs/uuid"
	"time"
)

// GetAllCustomers returns a slice of all customers
func (dbRepo *postgresDBRepo) GetAllCustomers() ([]models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var customers []models.Customer
	stmt := `
			select id ,first_name, last_name, birthdate, email, gender, created_at, updated_at
			from customers`

	rows, err := dbRepo.DB.QueryContext(ctx, stmt)
	if err != nil {
		return customers, err
	}

	for rows.Next() {
		var c models.Customer
		err = rows.Scan(
			&c.Uuid,
			&c.FirstName,
			&c.LastName,
			&c.Birthdate,
			&c.Email,
			&c.Gender,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return customers, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// InsertCustomer insert a customer into the database
func (dbRepo *postgresDBRepo) InsertCustomer(c models.Customer) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	uuid, err := uuid.NewV4()
	fmt.Println(uuid)
	if err != nil {
		return "", err
	}
	stmt := `
			insert into customers
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

// GetCustomerByID returns one customer by id
func (dbRepo *postgresDBRepo) GetCustomerByID(uuid uuid.UUID) (models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var customer models.Customer
	stmt := `
		select id, first_name, last_name, birthdate, email, gender, created_at, updated_at
    	from customers
		where id = $1`

	row := dbRepo.DB.QueryRowContext(ctx, stmt,
		uuid,
	)
	err := row.Scan(
		&customer.Uuid,
		&customer.FirstName,
		&customer.LastName,
		&customer.Birthdate,
		&customer.Email,
		&customer.Gender,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

// UpdateCustomer updates a customer in the database
func (dbRepo *postgresDBRepo) UpdateCustomer(c models.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update customers set 
				first_name = $2,
				last_name = $3,
				birthdate = $4,
				email = $5,
				gender = $6, 
				updated_at = $7
			where id = $1`

	_, err := dbRepo.DB.ExecContext(ctx, query,
		c.Uuid,
		c.FirstName,
		c.LastName,
		c.Birthdate,
		c.Email,
		c.Gender,
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}
