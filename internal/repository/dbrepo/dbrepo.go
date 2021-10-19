package dbrepo

import (
	"database/sql"
	"github.com/ekateryna-tln/wallester_task/internal/config"
	"github.com/ekateryna-tln/wallester_task/internal/repository"
)

type postgresDBRepo struct {
	App *config.App
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.App) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
