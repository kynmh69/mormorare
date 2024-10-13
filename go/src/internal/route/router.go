package route

import (
	"github.com/kynmh69/mormorare/internal/app"
)

func User(router *app.Engine) {
	router.Engine.GET("/users")
	router.Engine.POST("/users")
	router.Engine.PUT("/users")
	router.Engine.DELETE("/users")
}
