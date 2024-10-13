package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Engine struct {
	engine *gin.Engine
	db     *gorm.DB
}

func NewEngine(db *gorm.DB) *Engine {
	engine := gin.Default()
	return &Engine{engine: engine, db: db}
}
