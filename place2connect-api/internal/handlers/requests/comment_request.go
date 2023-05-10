package handlers

import (
	"github.com/fxfrancky/place2connect-api/internal/models"
	"github.com/gofiber/fiber"
)

// Comment Request
type CommentRequest struct {
	UserID             uint   `json:"userID" validate:"required"`
	PostID             uint   `json:"postID" validate:"required"`
	CommentDescription string `json:"commentDescription" validate:"required,min=5"`
}

// Binding Comments Request
func BindCommentsRequest(c *fiber.Ctx, comment *models.Comment, commentRequest *CommentRequest) error {

	// Validate the commentRequest
	if err := c.BodyParser(commentRequest); err != nil {
		return err
	}
	// Map the comment
	comment.UserID = commentRequest.UserID
	comment.PostID = commentRequest.PostID
	comment.CommentDescription = commentRequest.CommentDescription

	return nil
}

// Populate CommentsRequest
func populateCommentsRequest(comment *models.Comment, commentRequest *CommentRequest) {
	commentRequest.UserID = comment.UserID
	commentRequest.PostID = comment.PostID
	commentRequest.CommentDescription = comment.CommentDescription
}
