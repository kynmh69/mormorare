package route

import (
	"github.com/kynmh69/mormorare/internal/app"
)

func User(router *app.Engine) {
	router.Engine.GET("/users", router.User.Retrieve)
	router.Engine.POST("/users", router.User.Create)
	router.Engine.PUT("/users", router.User.Update)
	router.Engine.DELETE("/users", router.User.Delete)
}
