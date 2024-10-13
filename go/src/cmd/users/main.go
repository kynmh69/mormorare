package main

import (
	_ "github.com/kynmh69/mormorare/docs"
	"github.com/kynmh69/mormorare/internal/app"
	"github.com/kynmh69/mormorare/internal/database"
	"github.com/kynmh69/mormorare/internal/domain"
	"github.com/kynmh69/mormorare/pkg/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	// initialize the application
	logging.Initialize()
}

// @title Mormorare API
// @version 1.0
// license.name MIT
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
	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
