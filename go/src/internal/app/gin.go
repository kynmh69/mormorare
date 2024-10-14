package app

import (
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/mormorare/internal/domain/repository"
	"github.com/kynmh69/mormorare/internal/handler"
	"github.com/kynmh69/mormorare/pkg/logging"
	"gorm.io/gorm"
	"time"
)

type Engine struct {
	Engine *gin.Engine
	Db     *gorm.DB
	Repo   *repository.UserRepository
	User   *handler.UserHandler
	api    *gin.RouterGroup
}

func NewEngine(db *gorm.DB, repo *repository.UserRepository, userHandler *handler.UserHandler) *Engine {
	engine := gin.Default()
	engine.Use(ginZap.Ginzap(logging.GetZapLogger(), time.RFC3339, true))
	engine.Use(ginZap.RecoveryWithZap(logging.GetZapLogger(), true))
	return &Engine{
		Engine: engine,
		Db:     db,
		Repo:   repo,
		User:   userHandler,
	}
}

func (e *Engine) Run() {
	logger := logging.GetLogger()
	err := e.Engine.Run()
	if err != nil {
		logger.Panicln(err)
	}
}

func (e *Engine) Route() {
	e.api = e.Engine.Group("/api/v1")
	e.routeUser()
}

func (e *Engine) routeUser() {
	e.api.GET("/users", e.User.Retrieve)
	e.api.POST("/users", e.User.Create)
	e.api.PUT("/users/:id", e.User.Update)
	e.api.DELETE("/users/:username", e.User.Delete)
}
