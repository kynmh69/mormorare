package main

import (
	_ "github.com/kynmh69/mormorare/docs"
	"github.com/kynmh69/mormorare/internal/domain"
	"github.com/kynmh69/mormorare/internal/wire"
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
	// initialize the application
	r := wire.InitializeEngine()
	// initialize the database
	if err := r.Db.AutoMigrate(&domain.User{}); err != nil {
		logging.GetLogger().Panicln(err)
	} else {
		logging.GetLogger().Info("Database migrated")
	}
	r.Route()
	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
