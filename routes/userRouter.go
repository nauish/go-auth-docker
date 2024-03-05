package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/nauish/go-auth-docker/controllers"
	middleware "github.com/nauish/go-auth-docker/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users/me", controller.GetUser())
}
