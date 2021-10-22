package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/ekateryna-tln/wallester_task/internal/config"
	"github.com/ekateryna-tln/wallester_task/internal/driver"
	"github.com/ekateryna-tln/wallester_task/internal/handlers"
	"github.com/ekateryna-tln/wallester_task/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.App
var session *scs.SessionManager

// main is the main application function
func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

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

	// connect to database
	log.Println("connection to database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=wallester user=postgres password=Saule1234")
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
