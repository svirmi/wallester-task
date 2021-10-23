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
const portNumber = ":80"
const dbName = ""
const dbUser = ""
const dbPass = ""

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

func run() (*driver.DB, error) {

	app.UseCache = false
	app.CookieSecure = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.CookieSecure
	app.Session = session

	InitLocaleBundle()

	// connect to database
	log.Println("connection to database")
	db, err := driver.ConnectSQL(fmt.Sprintf("host=localhost port=5432 dbname=%s user=%s password=%s", dbName, dbUser, dbPass))
	if err != nil {
		log.Fatal("cannot connect to database:", err)
		return nil, err
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("create template cache error:", err)
		return nil, err
	}
	app.TemplateCache = tc
	app.UseCache = false

	render.SetRenderApp(&app)
	handlers.SetHandlersRepo(handlers.NewRepo(&app, db))

	return db, nil
}
