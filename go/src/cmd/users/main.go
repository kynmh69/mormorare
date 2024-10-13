package main

import (
	"github.com/kynmh69/mormorare/internal/app"
	"github.com/kynmh69/mormorare/internal/database"
	"github.com/kynmh69/mormorare/pkg/logging"
)

func init() {
	// initialize the application
	logging.Initialize()
}

func main() {
	// initialize the database
	db := database.NewPostgres()
	// initialize the application
	r := app.NewEngine(db)
	r.Route()
	r.Run()
}
