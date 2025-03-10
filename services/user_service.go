package services

import (
	"errors"
	"payment-system/config"
	"payment-system/models"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser hashes password and creates a new user
func RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)
	return config.DB.Create(user).Error
}

// AuthenticateUser verifies login credentials
func AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Compare stored hash with provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}
