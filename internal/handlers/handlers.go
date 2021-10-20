package handlers

import (
	"fmt"
	"github.com/ekateryna-tln/wallester_task/internal/config"
	"github.com/ekateryna-tln/wallester_task/internal/driver"
	"github.com/ekateryna-tln/wallester_task/internal/forms"
	"github.com/ekateryna-tln/wallester_task/internal/models"
	"github.com/ekateryna-tln/wallester_task/internal/render"
	"github.com/ekateryna-tln/wallester_task/internal/repository"
	"github.com/ekateryna-tln/wallester_task/internal/repository/dbrepo"
	"log"
	"net/http"
)

// Repo the repository used by handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.App
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(app *config.App, db *driver.DB) *Repository {
	return &Repository{
		App: app,
		DB:  dbrepo.NewPostgresRepo(db.SQL, app),
	}
}

// SetHandlersRepo set the repository for the handlers
func SetHandlersRepo(r *Repository) {
	Repo = r
}

// ShowAllCustomers shows all customers
func (repository *Repository) ShowAllCustomers(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "all-customers.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Fatal("can not render template:", err)
	}
}

// ShowCustomer renders the customer page
func (repository *Repository) ShowCustomer(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "all-customers.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Fatal("can not render template:", err)
	}
}

// ShowCustomerForm renders the add customer page and display form
func (repository *Repository) ShowCustomerForm(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "customers-form.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
	if err != nil {
		log.Fatal("can not render template:", err)
	}
}

// AddCustomer handles the posting of a customer from
func (repository *Repository) AddCustomer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("can not parse form:", err)
		repository.App.Session.Put(r.Context(), "error", "Something went wrong. Please contact to customer support.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	form := forms.New(r.PostForm)
	form.CheckRequiredFields("first_name", "last_name", "email", "birthdate")
	form.IsEmail("email")
	form.MaxLength("first_name", 100)
	form.MaxLength("last_name", 100)
	form.IsValidBirthdate("birthdate", 18, 60)

	customer := models.Customer{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Birthdate: r.Form.Get("birthdate"),
		Email:     r.Form.Get("email"),
		Gender:    r.Form.Get("gender"),
	}

	if !form.Valid() {
		data := make(map[string]interface{})
		data["customer"] = customer
		render.Template(w, r, "customers-form.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	uuid, err := repository.DB.InsertCustomer(customer)
	if err != nil {
		log.Println("can not insert a new customer to the database, ", err)
		repository.App.Session.Put(r.Context(), "error", "Can not add a new customer")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	repository.App.Session.Put(r.Context(), "flash", fmt.Sprintf("Customer created successfully. Customer identifier: %s", uuid))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// EditCustomer renders the edit customer page and display form
func (repository *Repository) EditCustomer(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "customers-form.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Fatal("can not render template:", err)
	}
}
