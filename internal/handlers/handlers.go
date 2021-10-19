package handlers

import (
	"github.com/ekateryna-tln/wallester_task/internal/config"
	"github.com/ekateryna-tln/wallester_task/internal/driver"
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

// AddCustomer renders the add customer page and display form
func (repository *Repository) AddCustomer(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "add-customers.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Fatal("can not render template:", err)
	}
}

// EditCustomer renders the edit customer page and display form
func (repository *Repository) EditCustomer(w http.ResponseWriter, r *http.Request) {
	err := render.Template(w, r, "add-customers.page.tmpl", &models.TemplateData{})
	if err != nil {
		log.Fatal("can not render template:", err)
	}
}
