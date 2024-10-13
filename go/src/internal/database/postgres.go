package database

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/mormorare/configs"
	"github.com/kynmh69/mormorare/consts"
	"github.com/kynmh69/mormorare/pkg/env"
	"github.com/kynmh69/mormorare/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewPostgres() *gorm.DB {
	var err error
	logger := logging.GetLogger()
	// Connect to Postgres
	config := getConfig()
	// Connect to Postgres
	logger.Debug(config)
	db, err = gorm.Open(postgres.Open(config.FormatDsn()), &gorm.Config{})
	if err != nil {
		logger.Panic(err)
	}
	return GetDB()
}

func GetDB() *gorm.DB {
	switch gin.Mode() {
	case gin.DebugMode:
		return db.Debug()
	case gin.TestMode, gin.ReleaseMode:
		return db
	default:
		return db
	}
}

func getHost() string {
	// Get the host from the environment
	return env.FindEnv(consts.PostgresHost)
}

func getPort() string {
	// Get the port from the environment
	return env.FindEnv(consts.PostgresPort)
}

func getUser() string {
	// Get the user from the environment
	return env.FindEnv(consts.PostgresUser)
}

func getPass() string {
	// Get the password from the environment
	return env.FindEnv(consts.PostgresPass)
}

func getDB() string {
	// Get the database from the environment
	return env.FindEnv(consts.PostgresDB)
}

func getTimeZone() string {
	// Get the timezone from the environment
	return env.FindEnv(consts.TimeZone)
}

func getConfig() *configs.PsqlConfig {
	return &configs.PsqlConfig{
		Host:     getHost(),
		Port:     getPort(),
		User:     getUser(),
		Password: getPass(),
		Dbname:   getDB(),
		TimeZone: getTimeZone(),
	}
}
