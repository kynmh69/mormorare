// +wireinject
package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kynmh69/mormorare/internal/app"
	"github.com/kynmh69/mormorare/internal/domain/repository"
	"github.com/kynmh69/mormorare/internal/handler"
	"gorm.io/gorm"
)

func InitializeUserHandler(repo repository.UserRepository, db *gorm.DB) *gin.Engine {
	wire.Build(app.NewEngine, handler.NewUserHandler, repository.NewUserRepository)
	return nil
}
