package handlers

import (
	"fmt"
	"github.com/ekateryna-tln/wallester_task/internal/config"
	"github.com/ekateryna-tln/wallester_task/internal/driver"
	"github.com/ekateryna-tln/wallester_task/internal/enums"
	"github.com/ekateryna-tln/wallester_task/internal/forms"
	"github.com/ekateryna-tln/wallester_task/internal/models"
	"github.com/ekateryna-tln/wallester_task/internal/render"
	"github.com/ekateryna-tln/wallester_task/internal/repository"
	"github.com/ekateryna-tln/wallester_task/internal/repository/dbrepo"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"net/url"
	"time"
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

// ShowHomePage renders the home page
func (repository *Repository) ShowHomePage(w http.ResponseWriter, r *http.Request) {
	if err := render.Template(w, r, "home.page.tmpl", &models.TemplateData{}); err != nil {
		log.Fatal("can not render template:", err)
	}
}

// ShowAllCustomers shows all customers
func (repository *Repository) ShowAllCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer

	customers, err := repository.DB.GetAllCustomers()
	if err != nil {
		log.Println("can not get customers from the database, ", err)
		repository.App.Session.Put(r.Context(), "error", "Could not get customers")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["customers"] = customers
	if err = render.Template(w, r, "all-customers.page.tmpl", &models.TemplateData{Data: data}); err != nil {
		log.Fatal("can not render template:", err)
	}
}

// ShowCustomer renders the customer page
func (repository *Repository) ShowCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u, err := uuid.FromString(id)
	if err != nil {
		log.Println("wrong uuid was given:", err)
		repository.App.Session.Put(r.Context(), "error", "Can not find the customer")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	customer, err = repository.DB.GetCustomerByID(u)
	if err != nil {
		log.Println("can not get the customer from the database, ", err)
		repository.App.Session.Put(r.Context(), "error", "Can not find the customer")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["customer"] = customer

	if err = render.Template(w, r, "customer.page.tmpl", &models.TemplateData{Data: data}); err != nil {
		log.Fatal("can not render template:", err)
	}
	return
}

// ShowCustomerForm renders the add customer page and display form
func (repository *Repository) ShowCustomerForm(w http.ResponseWriter, r *http.Request) {

	var customer models.Customer
	id := chi.URLParam(r, "id")
	if id != "" {
		u, err := uuid.FromString(id)
		if err != nil {
			log.Println("wrong uuid was given:", err)
			repository.App.Session.Put(r.Context(), "error", "Can not edit the customer")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		customer, err = repository.DB.GetCustomerByID(u)
		if err != nil {
			log.Println("can not get the customer from the database, ", err)
			repository.App.Session.Put(r.Context(), "error", "Can not find the customer")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	data := make(map[string]interface{})
	data["customer"] = customer
	data["minDate"] = time.Now().AddDate(-61, 0, +1)
	data["maxDate"] = time.Now().AddDate(-18, 0, 0)
	data["genderMale"] = enums.Male.String()
	data["genderFemale"] = enums.Female.String()

	if err := render.Template(w, r, "customers-form.page.tmpl", &models.TemplateData{Form: forms.New(nil), Data: data}); err != nil {
		log.Fatal("can not render template:", err)
	}
}

func (repository *Repository) getCustomerFromIncomingForm(formData url.Values) (models.Customer, *forms.Form) {

	customer := models.Customer{
		FirstName: formData.Get("first_name"),
		LastName:  formData.Get("last_name"),
		Email:     formData.Get("email"),
		Gender:    formData.Get("gender"),
	}

	form := forms.New(formData)
	form.CheckRequiredFields("first_name", "last_name", "email", "birthdate")
	form.IsEmail("email")
	form.MaxLength("first_name", 100)
	form.MaxLength("last_name", 100)
	birthdate, ok := form.IsValidDate("birthdate")
	if ok {
		form.IsValidAge("birthdate", birthdate, 18, 60)
		customer.Birthdate = birthdate
	}
	form.IsValidGender("gender")

	return customer, form
}

func (repository *Repository) redirectIfInvalidForm(w http.ResponseWriter, r *http.Request, customer models.Customer, form *forms.Form) {
	data := make(map[string]interface{})
	data["customer"] = customer
	data["minDate"] = time.Now().AddDate(-60, 0, -1)
	data["maxDate"] = time.Now().AddDate(-18, 0, +1)
	data["genderMale"] = enums.Male.String()
	data["genderFemale"] = enums.Female.String()
	if err := render.Template(w, r, "customers-form.page.tmpl", &models.TemplateData{Form: form, Data: data}); err != nil {
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

	customer, form := repository.getCustomerFromIncomingForm(r.PostForm)

	if !form.Valid() {
		repository.redirectIfInvalidForm(w, r, customer, form)
		return
	}

	u, err := repository.DB.InsertCustomer(customer)
	if err != nil {
		log.Println("can not insert a new customer to the database, ", err)
		repository.App.Session.Put(r.Context(), "error", "Can not add a new customer")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	repository.App.Session.Put(r.Context(), "flash", fmt.Sprintf("Customer created successfully. Customer identifier: %s", u))
	http.Redirect(w, r, "/customers", http.StatusSeeOther)
}

// EditCustomer renders the edit customer page and display form
func (repository *Repository) EditCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	var u uuid.UUID
	var err error
	id := chi.URLParam(r, "id")
	if id != "" {
		u, err = uuid.FromString(id)
		if err != nil {
			log.Println("wrong uuid was given:", err)
			repository.App.Session.Put(r.Context(), "error", "Can not edit the customer")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	err = r.ParseForm()
	if err != nil {
		log.Println("can not parse form:", err)
		repository.App.Session.Put(r.Context(), "error", "Something went wrong. Please contact to customer support.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	customer, form := repository.getCustomerFromIncomingForm(r.PostForm)
	customer.Uuid = u.String()

	if !form.Valid() {
		repository.redirectIfInvalidForm(w, r, customer, form)
		return
	}
	err = repository.DB.UpdateCustomer(customer)
	if err != nil {
		log.Println("can not update a customer:", err)
		repository.App.Session.Put(r.Context(), "error", "Can not add a new customer")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	repository.App.Session.Put(r.Context(), "flash", fmt.Sprintf("Customer updated successfully. Customer identifier: %s", u))
	http.Redirect(w, r, "/customers", http.StatusSeeOther)
}
