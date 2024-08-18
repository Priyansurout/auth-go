package user

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Age      int    `gorm:"not null"`
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// InitializeDB initializes the database connection and migrates the schema.
func InitializeDB(dns string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {

		return nil, err
	}

	// AutoMigrate will create or update the table based on the model.
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, err
	}

	return db, nil

}
