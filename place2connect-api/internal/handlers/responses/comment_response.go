package handlers

import (
	"time"

	"github.com/fxfrancky/place2connect-api/internal/models"
)

// Comment Response
type CommentResponse struct {
	ID                 uint      `json:"id,omitempty"`
	UserID             uint      `json:"user_id,omitempty"`
	PostID             uint      `json:"post_id,omitempty"`
	CommentDescription string    `json:"comment_description,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// Comment List Response
type CommentListResponse struct {
	Comments      []*CommentResponse `json:"comments"`
	CommentsCount int64              `json:"commentsCount"`
}

func NewCommentResponse(comment *models.Comment) *CommentResponse {
	comr := new(CommentResponse)
	comr.ID = comment.ID
	comr.UserID = comment.UserID
	comr.PostID = comment.PostID
	comr.CommentDescription = comment.CommentDescription
	return comr
}

func NewCommentListResponse(comments []models.Comment, count int64) *CommentListResponse {
	c := new(CommentListResponse)
	c.Comments = make([]*CommentResponse, 0)
	for _, cs := range comments {
		cr := NewCommentResponse(&cs)
		c.Comments = append(c.Comments, cr)
	}
	c.CommentsCount = count
	return c
}
