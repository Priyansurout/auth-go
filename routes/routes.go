package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/priyansurout/auth-microservice/controllers" // Updated import path
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/register", controllers.SignupHandler(db))
	r.POST("/login", controllers.LoginHandler(db))
	r.POST("/logout", controllers.LogoutHandler(db))
	// Add other routes as needed
}
