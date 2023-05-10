package handlers

import (
	"time"

	"github.com/fxfrancky/place2connect-api/internal/models"
)

// Like Response
type LikeResponse struct {
	ID            uint      `json:"id,omitempty"`
	UserID        uint      `json:"user_id,omitempty"`
	PostID        uint      `json:"post_id,omitempty"`
	IsLikablePost bool      `json:"is_likable_post,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Like List Response
type LikeListResponse struct {
	Likes      []*LikeResponse `json:"likes"`
	LikesCount int64           `json:"likesCount"`
}

func NewLikeResponse(like *models.Like) *LikeResponse {
	lr := new(LikeResponse)
	lr.ID = like.ID
	lr.UserID = like.UserID
	lr.PostID = like.PostID
	lr.IsLikablePost = like.IsLikablePost
	return lr
}

func NewLikeListResponse(likes []models.Like, count int64) *LikeListResponse {
	l := new(LikeListResponse)
	l.Likes = make([]*LikeResponse, 0)
	for _, ls := range likes {
		lr := NewLikeResponse(&ls)
		l.Likes = append(l.Likes, lr)
	}
	l.LikesCount = count
	return l
}
