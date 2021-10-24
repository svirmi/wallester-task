package main

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/ekateryna-tln/wallester_task/internal/config"
	"github.com/ekateryna-tln/wallester_task/internal/driver"
	"github.com/ekateryna-tln/wallester_task/internal/handlers"
	"github.com/ekateryna-tln/wallester_task/internal/render"
	"log"
	"net/http"
	"time"
)

// Parameters should be set according to personal settings
const portNumber = ":8080"
const dbName = ""
const dbUser = ""
const dbPass = ""
const dbHost = ""
const dbPort = ""

var app config.App
var session *scs.SessionManager

// main is the main application function
func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db.SQL)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	log.Fatal(srv.ListenAndServe())
}

// run adds app settings, connects to the database and sets template cache for the application
func run() (*driver.DB, error) {

	app.UseCache = false
	app.CookieSecure = false
	app.Session = serSession()

	InitLocaleBundle()

	// connect to database
	db, err := driver.ConnectSQL(fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", dbHost, dbPort, dbName, dbUser, dbPass))
	if err != nil {
		return nil, err
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		return nil, err
	}
	app.TemplateCache = tc
	app.UseCache = false

	render.SetRenderApp(&app)
	handlers.SetHandlersRepo(handlers.NewRepo(&app, db))

	return db, nil
}

// serSession sets session parameters
func serSession() *scs.SessionManager {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.CookieSecure
	return session
}
