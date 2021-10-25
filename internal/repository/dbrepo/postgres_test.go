package dbrepo

import (
	"errors"
	"github.com/gofrs/uuid"
	"testing"
	"time"

	"github.com/ekateryna-tln/wallester-task/internal/models"
)

func Test_postgresDBRepo_InsertCustomer(t *testing.T) {

	if err := dbRepo.TruncateCustomer(); err != nil {
		t.Errorf("TruncateCustomer error = %s", err)
		return
	}
	birthdateStr := time.Now().Format("2006-01-02")
	layout := "2006-01-02"
	birthdate, _ := time.Parse(layout, birthdateStr)

	customer := models.Customer{
		FirstName:   "TestFirstName",
		LastName:    "TestLastName",
		Birthdate:   birthdate,
		Email:       "TestCustomer@test.test",
		Gender:      "Male",
		SearchField: "testfirstname_testlastname",
	}

	type args struct {
		c models.Customer
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{"success_insert", args{c: customer}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUuid, err := dbRepo.InsertCustomer(tt.args.c)
			u, err := uuid.FromString(gotUuid)
			if err != nil {
				t.Errorf("InsertCustomer() error = %s, want %s", err, tt.wantErr)
				return
			}
			gotCustomer, err := dbRepo.GetCustomerByID(u)
			if err != nil {
				t.Errorf("GetCustomerByID() error = %s, want %s", err, tt.wantErr)
				return
			}
			if !tt.args.c.EqualBase(gotCustomer) {
				t.Errorf("InsertCustomer() got = %v, want %v", gotCustomer, tt.args.c)
			}
		})
	}
}

func Test_postgresDBRepo_GetAllCustomers(t *testing.T) {
	if err := dbRepo.TruncateCustomer(); err != nil {
		t.Errorf("TruncateCustomer error = %s", err)
		return
	}

	customerCount := 10
	birthdateStr := time.Now().Format("2006-01-02")
	layout := "2006-01-02"
	birthdate, _ := time.Parse(layout, birthdateStr)

	customer := models.Customer{
		FirstName:   "TestFirstName",
		LastName:    "TestLastName",
		Birthdate:   birthdate,
		Email:       "TestCustomer@test.test",
		Gender:      "Male",
		SearchField: "testfirstname_testlastname",
	}

	type args struct {
		c models.Customer
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{"valid_test", args{c: customer}, customerCount, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 10; i++ {
				_, _ = dbRepo.InsertCustomer(tt.args.c)
			}
			got, err := dbRepo.GetAllCustomers()
			if err != nil {
				t.Errorf("GetAllCustomers() error = %s, wantErr %s", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("GetAllCustomers() wrang customer emount. got = %d, want %d", len(got), tt.want)
			}
		})
	}
}

func Test_postgresDBRepo_UpdateCustomer(t *testing.T) {
	if err := dbRepo.TruncateCustomer(); err != nil {
		t.Errorf("TruncateCustomer error = %s", err)
		return
	}
	birthdateStr := time.Now().Format("2006-01-02")
	layout := "2006-01-02"
	birthdate, _ := time.Parse(layout, birthdateStr)

	customer := models.Customer{
		FirstName:   "TestFirstName",
		LastName:    "TestLastName",
		Birthdate:   birthdate,
		Email:       "TestCustomer@test.test",
		Gender:      "Male",
		SearchField: "testfirstname_testlastname",
	}

	newCustomer := models.Customer{
		FirstName:   "NewTestFirstName",
		LastName:    "NewTestLastName",
		Birthdate:   birthdate,
		Email:       "NewTestCustomer@test.test",
		Gender:      "Female",
		SearchField: "newtestfirstname_newtestlastname",
	}

	type args struct {
		c    models.Customer
		cNew models.Customer
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{"success_insert", args{c: customer, cNew: newCustomer}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUuid, err := dbRepo.InsertCustomer(tt.args.c)
			u, err := uuid.FromString(gotUuid)
			if err != nil {
				t.Errorf("InsertCustomer() error = %s, want %s", err, tt.wantErr)
				return
			}
			tt.args.cNew.Uuid = gotUuid
			err = dbRepo.UpdateCustomer(tt.args.cNew)
			gotCustomer, err := dbRepo.GetCustomerByID(u)
			if err != nil {
				t.Errorf("GetCustomerByID() error = %s, want %s", err, tt.wantErr)
				return
			}
			if !tt.args.cNew.EqualBase(gotCustomer) {
				t.Errorf("UpdateCustomer() got = %v, want %v", gotCustomer, tt.args.cNew)
			}
		})
	}
}

func Test_postgresDBRepo_SearchCustomers(t *testing.T) {
	if err := dbRepo.TruncateCustomer(); err != nil {
		t.Errorf("TruncateCustomer error = %s", err)
		return
	}
	customerCount := 1
	birthdateStr := time.Now().Format("2006-01-02")
	layout := "2006-01-02"
	birthdate, _ := time.Parse(layout, birthdateStr)

	customer := models.Customer{
		FirstName:   "TestFirstName",
		LastName:    "TestLastName",
		Birthdate:   birthdate,
		Email:       "TestCustomer@test.test",
		Gender:      "Male",
		SearchField: "testfirstname_testlastname",
	}

	newCustomer := models.Customer{
		FirstName:   "NewTestFirstName",
		LastName:    "NewTestLastName",
		Birthdate:   birthdate,
		Email:       "NewTestCustomer@test.test",
		Gender:      "Female",
		SearchField: "newtestfirstname_newtestlastname",
	}

	type args struct {
		c      models.Customer
		cNew   models.Customer
		search string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{"success_search", args{c: customer, cNew: newCustomer, search: "new"}, customerCount, nil},
		{"invalid_empty_search", args{c: customer, cNew: newCustomer, search: ""}, 0, errors.New("no data to search")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUuid, err := dbRepo.InsertCustomer(tt.args.c)
			_, err = uuid.FromString(gotUuid)
			if err != nil {
				t.Errorf("InsertCustomer() error = %s, want %s", err, tt.wantErr)
				return
			}
			gotUuid, err = dbRepo.InsertCustomer(tt.args.cNew)
			_, err = uuid.FromString(gotUuid)
			if err != nil {
				t.Errorf("InsertCustomer() error = %s, want %s", err, tt.wantErr)
				return
			}

			gotCustomers, err := dbRepo.SearchCustomers(tt.args.search)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("SearchCustomers() error = %s, want %s", err, tt.wantErr.Error())
				return
			}
			if len(gotCustomers) != tt.want {
				t.Errorf("SearchCustomers() got = %d, want %d", len(gotCustomers), tt.want)
			}
		})
	}
}
