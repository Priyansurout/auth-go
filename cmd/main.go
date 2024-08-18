package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/priyansurout/auth-microservice/internal/user"
	"github.com/priyansurout/auth-microservice/routes"
)

func main() {

	fmt.Println("Hello, world")

	// DSN: Data Source Name
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/authdb?charset=utf8&parseTime=True&loc=Local"

	// Initialize the database
	db, err := user.InitializeDB(dsn)

	if err != nil {
		fmt.Println("Fail to connect!")
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if db != nil {
		fmt.Println("GOOD!")
	}

	// Initialize Gin router
	router := gin.Default()

	// Load the routes
	routes.AuthRoutes(router, db)

	// Run the server
	router.Run(":8000")
}
