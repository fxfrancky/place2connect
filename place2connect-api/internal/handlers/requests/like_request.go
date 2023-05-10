package handlers

import (
	"github.com/fxfrancky/place2connect-api/internal/models"
	"github.com/gofiber/fiber"
)

// Likes Request
type LikeRequest struct {
	UserID        uint `json:"userID" validate:"required"`
	PostID        uint `json:"postID" validate:"required"`
	IsLikablePost bool `json:"isLikablePost" validate:"required,boolean"`
}

// Binding Likes Request
func BindLikeRequest(c *fiber.Ctx, like *models.Like, likeRequest *LikeRequest) error {

	// Validate the likeRequest
	if err := c.BodyParser(likeRequest); err != nil {
		return err
	}
	// Map the like
	like.IsLikablePost = likeRequest.IsLikablePost
	like.UserID = likeRequest.UserID
	like.PostID = likeRequest.PostID

	return nil
}

// Populate Likes Request
func PopulateLikeRequest(like *models.Like, likeRequest *LikeRequest) {
	likeRequest.IsLikablePost = like.IsLikablePost
	likeRequest.PostID = like.PostID
	likeRequest.UserID = like.UserID
}
