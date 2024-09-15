package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/GreenTeaProgrammers/MatsuriSNS/ent"

	"github.com/gin-gonic/gin"
)

// PostController handles the post-related routes.
type PostController struct {
	client *ent.Client
}

// NewPostController creates a new PostController.
func NewPostController(client *ent.Client) *PostController {
	return &PostController{client}
}

// CreatePost handles the creation of a new post.
func (pc *PostController) CreatePost(c *gin.Context) {
	var input struct {
		UserID    int     `json:"user_id" binding:"required"`
		EventID   int     `json:"event_id" binding:"required"`
		Content   string  `json:"content" binding:"required"`
		LocationX float64 `json:"location_x" binding:"required"`
		LocationY float64 `json:"location_y" binding:"required"`
		VideoURL  string  `json:"video_url"`
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	// Verify that the User exists
	_, err := pc.client.User.Get(ctx, input.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Verify that the Event exists
	_, err = pc.client.Event.Get(ctx, input.EventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event not found"})
		return
	}

	// Create the post in the database
	post, err := pc.client.Post.Create().
		SetUserID(input.UserID).
		SetEventID(input.EventID).
		SetContent(input.Content).
		SetLocationX(input.LocationX).
		SetLocationY(input.LocationY).
		SetNillableVideoURL(&input.VideoURL).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPost retrieves a post by its ID.
func (pc *PostController) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := pc.client.Post.Get(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// ListPosts retrieves all posts.
func (pc *PostController) ListPosts(c *gin.Context) {
	posts, err := pc.client.Post.Query().All(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// UpdatePost updates an existing post.
func (pc *PostController) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var input struct {
		Content   string  `json:"content"`
		LocationX float64 `json:"location_x"`
		LocationY float64 `json:"location_y"`
		VideoURL  string  `json:"video_url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := pc.client.Post.Get(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Update the post
	post, err = post.Update().
		SetContent(input.Content).
		SetLocationX(input.LocationX).
		SetLocationY(input.LocationY).
		SetNillableVideoURL(&input.VideoURL).
		SetUpdatedAt(time.Now()).
		Save(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost deletes an existing post.
func (pc *PostController) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	err = pc.client.Post.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
