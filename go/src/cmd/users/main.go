package main

import (
	"github.com/kynmh69/mormorare/internal/app"
	"github.com/kynmh69/mormorare/internal/database"
	"github.com/kynmh69/mormorare/internal/domain"
	"github.com/kynmh69/mormorare/pkg/logging"
)

func init() {
	// initialize the application
	logging.Initialize()
}

func main() {
	// initialize the database
	db := database.NewPostgres()
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		logging.GetLogger().Panicln(err)
	} else {
		logging.GetLogger().Info("Database migrated")
	}
	// initialize the application
	r := app.NewEngine(db)
	r.Route()
	r.Run()
}
