package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/mormorare/internal/handler"
	"github.com/kynmh69/mormorare/internal/route"
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
	route.User(e)
}

// createUserHandler Create User Handler
func (e *Engine) createUserHandler() {
	e.User = handler.NewUserHandler(e.DB)
}
