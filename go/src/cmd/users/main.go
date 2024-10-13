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

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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
