package handlers

import (
	"log"

	"github.com/fxfrancky/place2connect-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Post Request
type PostRequest struct {
	UserID uint `json:"userID" validate:"required"`
	// FirstName       string `json:"first_name" validate:"required,min=3"`
	// LastName        string `json:"last_name" validate:"required,min=3"`
	// Location        string `json:"location,omitempty"`
	Description string `json:"description" validate:"required,min=8"`
	PicturePath string `json:"picturePath,omitempty"`
	// UserPicturePath string `json:"user_picture_path,omitempty"`
	// LikesRequestList   []LikeRequest
	// CommentRequestList []CommentRequest
}

// UserFriendRequest Request
type AddUserFriendRequest struct {
	UserID   uint `json:"userID" validate:"required"`
	FriendID uint `json:"friendID" validate:"required"`
}

func BindUserFriendRequest(c *fiber.Ctx, addUserFriendRequest *AddUserFriendRequest) error {
	// Validate the userFriendRequest
	if err := c.BodyParser(addUserFriendRequest); err != nil {
		return err
	}
	return nil
}

// Binding Posts Request
func BindPostsRequest(c *fiber.Ctx, post *models.Post, postRequest *PostRequest) error {

	// Validate the postRequest
	if err := c.BodyParser(postRequest); err != nil {
		return err
	}
	log.Println("Binding Post Request ", postRequest.UserID)
	// Map the post
	post.UserID = postRequest.UserID
	// post.FirstName = postRequest.FirstName
	// post.LastName = postRequest.LastName
	// post.Location = postRequest.Location
	post.Description = postRequest.Description
	post.PicturePath = postRequest.PicturePath
	// post.UserPicturePath = postRequest.UserPicturePath
	// post.Likes()
	// post.Comments()

	return nil
}

// Populate Posts Request
func PopulatePostsRequest(post *models.Post, postRequest *PostRequest) {
	postRequest.UserID = post.UserID
	// postRequest.FirstName = post.FirstName
	// postRequest.LastName = post.LastName
	// postRequest.Location = post.Location
	postRequest.Description = post.Description
	postRequest.PicturePath = post.PicturePath
	// postRequest.UserPicturePath = post.UserPicturePath
}
