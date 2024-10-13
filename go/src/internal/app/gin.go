package app

import (
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/mormorare/internal/handler"
	"github.com/kynmh69/mormorare/pkg/logging"
	"gorm.io/gorm"
	"time"
)

type Engine struct {
	Engine *gin.Engine
	DB     *gorm.DB
	User   *handler.UserHandler
	api    *gin.RouterGroup
}

func NewEngine(db *gorm.DB) *Engine {
	engine := gin.Default()
	engine.Use(ginZap.Ginzap(logging.GetZapLogger(), time.RFC3339, true))
	engine.Use(ginZap.RecoveryWithZap(logging.GetZapLogger(), true))
	return &Engine{Engine: engine, DB: db}
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
	e.User = handler.NewUserHandler(e.DB)
}

func (e *Engine) routeUser() {
	e.api.GET("/users", e.User.Retrieve)
	e.api.POST("/users", e.User.Create)
	e.api.PUT("/users/:id", e.User.Update)
	e.api.DELETE("/users/:username", e.User.Delete)
}
