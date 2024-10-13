package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
