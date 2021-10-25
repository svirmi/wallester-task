package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/ekateryna-tln/wallester-task/internal/models"
	"github.com/gofrs/uuid"
)

// GetAllCustomers returns a slice of all customers
func (dbRepo *postgresDBRepo) GetAllCustomers() ([]models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var customers []models.Customer
	stmt := `
			select id ,first_name, last_name, birthdate, email, gender, created_at, updated_at
			from customers
			order by updated_at desc`

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

// SearchCustomers returns a slice of found customers
func (dbRepo *postgresDBRepo) SearchCustomers(searchExpression string) ([]models.Customer, error) {
	var customers []models.Customer
	if searchExpression == "" {
		return customers, errors.New("no data to search")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
			select id, first_name, last_name, birthdate, email, gender, created_at, updated_at
			from customers where search_field like '%' || $1 || '%'`

	rows, err := dbRepo.DB.QueryContext(ctx, stmt, searchExpression)
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

	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	stmt := `
			insert into customers
			(id ,first_name, last_name, birthdate, email, gender, search_field, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = dbRepo.DB.ExecContext(ctx, stmt,
		u,
		c.FirstName,
		c.LastName,
		c.Birthdate,
		c.Email,
		c.Gender,
		c.SearchField,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return "", err
	}

	return u.String(), nil
}

// GetCustomerByID returns one customer by id
func (dbRepo *postgresDBRepo) GetCustomerByID(u uuid.UUID) (models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var customer models.Customer
	stmt := `
		select id, first_name, last_name, birthdate, email, gender, created_at, updated_at
    	from customers
		where id = $1`

	row := dbRepo.DB.QueryRowContext(ctx, stmt,
		u,
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
				search_field = $7, 
				updated_at = $8
			where id = $1`

	_, err := dbRepo.DB.ExecContext(ctx, query,
		c.Uuid,
		c.FirstName,
		c.LastName,
		c.Birthdate,
		c.Email,
		c.Gender,
		c.SearchField,
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}

// TruncateCustomer removes all customers. Created for the test purpose.
// Should be used carefully.
func (dbRepo *postgresDBRepo) TruncateCustomer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `truncate customers`

	_, err := dbRepo.DB.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}
