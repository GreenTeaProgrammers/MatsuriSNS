package routes

import (
	"github.com/GreenTeaProgrammers/MatsuriSNS/controllers"
	"github.com/GreenTeaProgrammers/MatsuriSNS/ent"
	"github.com/gin-gonic/gin"
)

// PostRoutes defines the routes for post-related actions.
func SetupPostRoutes(router *gin.Engine, client *ent.Client) {
	postController := controllers.NewPostController(client)

	// Create a new post
	router.POST("/posts", postController.CreatePost)

	// Get a post by ID
	router.GET("/posts/:id", postController.GetPost)

	// List all posts
	router.GET("/posts", postController.ListPosts)

	// Update a post by ID
	router.PUT("/posts/:id", postController.UpdatePost)

	// Delete a post by ID
	router.DELETE("/posts/:id", postController.DeletePost)
}
