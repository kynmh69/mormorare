// +wireinject
package app

import (
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/mormorare/internal/domain/repository"
	"github.com/kynmh69/mormorare/internal/handler"
	"github.com/kynmh69/mormorare/pkg/logging"
	"time"
)

type Engine struct {
	Engine *gin.Engine
	Repo   repository.UserRepository
	User   *handler.UserHandler
	api    *gin.RouterGroup
}

func NewEngine(repo repository.UserRepository) *Engine {
	engine := gin.Default()
	engine.Use(ginZap.Ginzap(logging.GetZapLogger(), time.RFC3339, true))
	engine.Use(ginZap.RecoveryWithZap(logging.GetZapLogger(), true))
	return &Engine{Engine: engine, Repo: repo}
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
	e.createUserHandler()
	e.routeUser()
}

// createUserHandler Create User Handler
func (e *Engine) createUserHandler() {
	e.User = handler.NewUserHandler(e.Repo)
}

func (e *Engine) routeUser() {
	e.api.GET("/users", e.User.Retrieve)
	e.api.POST("/users", e.User.Create)
	e.api.PUT("/users/:id", e.User.Update)
	e.api.DELETE("/users/:username", e.User.Delete)
}
