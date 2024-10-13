package database

import (
	"github.com/kynmh69/mormorare/configs"
	"github.com/kynmh69/mormorare/consts"
	"github.com/kynmh69/mormorare/pkg/env"
	"github.com/kynmh69/mormorare/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewPostgres() {
	var err error
	logger := logging.GetLogger()
	// Connect to Postgres
	config := getConfig()
	// Connect to Postgres
	db, err = gorm.Open(postgres.Open(config.FormatDsn()), &gorm.Config{})
	if err != nil {
		logger.Panic(err)
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

func getConfig() *configs.PsqlConfig {
	return &configs.PsqlConfig{
		Host:     getHost(),
		Port:     getPort(),
		User:     getUser(),
		Password: getPass(),
		Dbname:   getDB(),
	}
}
