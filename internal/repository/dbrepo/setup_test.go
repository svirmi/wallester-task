package dbrepo

import (
	"fmt"
	"github.com/ekateryna-tln/wallester-task/internal/config"
	"github.com/ekateryna-tln/wallester-task/internal/driver"
	"log"
	"os"
	"testing"
)

const dbName = "wallester_tests"
const dbUser = "postgres"
const dbPass = "Saule1234"
const dbHost = "localhost"
const dbPort = "5432"

var db *driver.DB
var dbRepo *postgresDBRepo

// TestMain contains setup for database tests.
// Pay attention that tests must use separate database for tests
func TestMain(m *testing.M) {
	var err error
	defer closeDB()

	// connect to database
	db, err = driver.ConnectSQL(fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", dbHost, dbPort, dbName, dbUser, dbPass))
	if err != nil {
		log.Fatal("could not connect to the database for tests")
	}
	dbRepo = &postgresDBRepo{
		App: &config.App{},
		DB:  db.SQL,
	}

	os.Exit(m.Run())
}

func closeDB() {
	err := db.SQL.Close()
	if err != nil {
		return
	}
	fmt.Printf("Teardown completed")
}
