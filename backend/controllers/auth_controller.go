package controllers

import (
	"context"
	"net/http"

	"github.com/GreenTeaProgrammers/MatsuriSNS/ent"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent/user"
	"github.com/GreenTeaProgrammers/MatsuriSNS/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	client *ent.Client
}

// NewAuthController creates a new AuthController with an ent.Client.
func NewAuthController(client *ent.Client) *AuthController {
	return &AuthController{client: client}
}

// Register handles user registration
func (ac *AuthController) Register(c *gin.Context) {
	type RegisterInput struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user in the database
	user, err := ac.client.User.
		Create().
		SetUsername(input.Username).
		SetEmail(input.Email).
		SetPassword(hashedPassword).
		Save(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful", "user": user})
}

// Login handles user login
func (ac *AuthController) Login(c *gin.Context) {
	type LoginInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user by email
	user, err := ac.client.User.
		Query().
		Where(user.Email(input.Email)).
		Only(context.Background())
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the password matches
	if err := utils.CheckPassword(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Success
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

