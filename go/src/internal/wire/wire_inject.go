//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kynmh69/mormorare/internal/app"
	"github.com/kynmh69/mormorare/internal/database"
	"github.com/kynmh69/mormorare/internal/domain/repository"
	"github.com/kynmh69/mormorare/internal/handler"
)

func InitializeEngine() *app.Engine {
	wire.Build(
		database.NewPostgres,
		repository.NewUserRepository,
		handler.NewUserHandler,
		app.NewEngine,
	)
	return nil
}
