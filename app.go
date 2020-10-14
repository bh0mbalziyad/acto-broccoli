// app.go

package main

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	fmt.Println("Database started")
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {}
