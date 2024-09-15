package controllers

import (
	"context"
	"log"
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
		Password string `json:"password" binding:"required,min=6"`
	}

	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// エラーメッセージをログに出力
		log.Printf("Registration error: %v", err)
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
		SetHashedPassword(hashedPassword). // Use SetHashedPassword instead of SetPassword
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
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユーザー名でユーザーを検索
	user, err := ac.client.User.
		Query().
		Where(user.UsernameEQ(input.Username)).
		Only(context.Background())
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// パスワードの確認
	if err := utils.CheckPassword(user.HashedPassword, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// ログイン成功
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
