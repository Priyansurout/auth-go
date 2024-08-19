package controllers

import (
	"net/http"

	"github.com/priyansurout/auth-microservice/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/priyansurout/auth-microservice/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// func LoginHandler(c *gin.Context) {
// 	// Handle user login

//		c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
//	}
//
// Login handles user login requests
func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqUser user.User
		var foundUser user.User

		// Bind JSON to reqUser
		if err := c.ShouldBindJSON(&reqUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Find user by username
		if err := db.Where("username = ?", reqUser.Username).First(&foundUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// Check if the password is correct
		if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(reqUser.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// Generate JWT token
		token, err := utils.GenerateJWT(foundUser.ID, foundUser.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.SetCookie("auth_token", token, 24*3600, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
	}
}

func SignupHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Handle user registration

		var reqUser user.User

		if err := c.ShouldBindJSON(&reqUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existingUser user.User
		if err := db.Where("username = ?", reqUser.Username).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		reqUser.Password = string(hashedPassword)

		if err := db.Create(&reqUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		// Return a success response
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

	}
}

func LogoutHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Handle user logout
		c.JSON(http.StatusOK, gin.H{"message": "User logged out"})
	}
}
