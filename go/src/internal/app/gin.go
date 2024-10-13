package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/mormorare/internal/handler"
	"github.com/kynmh69/mormorare/pkg/logging"
	"gorm.io/gorm"
)

type Engine struct {
	Engine *gin.Engine
	DB     *gorm.DB
	User   *handler.UserHandler
}

func NewEngine(db *gorm.DB) *Engine {
	engine := gin.Default()
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
	e.createUserHandler()
	e.routeUser()
}

// createUserHandler Create User Handler
func (e *Engine) createUserHandler() {
	e.User = handler.NewUserHandler(e.DB)
}

func (e *Engine) routeUser() {
	e.Engine.GET("/users", e.User.Retrieve)
	e.Engine.POST("/users", e.User.Create)
	e.Engine.PUT("/users", e.User.Update)
	e.Engine.DELETE("/users", e.User.Delete)
}
