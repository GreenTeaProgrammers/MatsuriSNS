package routes

import (
	"github.com/GreenTeaProgrammers/MatsuriSNS/controllers"
	"github.com/gin-gonic/gin"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent"
)

// SetupAuthRoutes initializes authentication routes
func SetupAuthRoutes(r *gin.Engine, client *ent.Client) {
	authController := controllers.NewAuthController(client)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}
}

