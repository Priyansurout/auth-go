package routes

import (
	"github.com/gin-gonic/gin"
	auth "github.com/priyansurout/auth-microservice/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", auth.RegisterHandler)
	r.POST("/login", auth.LoginHandler)
	r.POST("/logout", auth.LogoutHandler)
	// Add other routes as needed
}
