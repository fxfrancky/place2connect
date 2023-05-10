package handlers

import (
	"time"

	"github.com/fxfrancky/place2connect-api/internal/models"
)

// Post Response
type PostResponse struct {
	ID              uint               `json:"id,omitempty"`
	UserID          uint               `json:"userID,omitempty"`
	FirstName       string             `json:"firstName,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	Location        string             `json:"location,omitempty"`
	Description     string             `json:"description,omitempty"`
	PicturePath     string             `json:"picturePath,omitempty"`
	UserPicturePath string             `json:"userPicturePath,omitempty"`
	Likes           []*LikeResponse    `json:"likes"`
	Comments        []*CommentResponse `json:"comments"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"`
}

// Post List Response
type PostListResponse struct {
	Posts     []*PostResponse `json:"posts"`
	PostCount int64           `json:"postCount"`
}

func NewPostResponse(post *models.Post) *PostResponse {
	postListLikes := NewLikeListResponse(post.Likes, int64(len(post.Likes)))
	postListComments := NewCommentListResponse(post.Comments, int64(len(post.Comments)))

	pr := new(PostResponse)
	pr.ID = post.ID
	pr.UserID = post.UserID
	pr.FirstName = post.FirstName
	pr.LastName = post.LastName
	pr.Location = post.Location
	pr.Description = post.Description
	pr.PicturePath = post.PicturePath
	pr.UserPicturePath = post.UserPicturePath
	pr.Likes = postListLikes.Likes
	pr.Comments = postListComments.Comments
	pr.CreatedAt = post.CreatedAt
	pr.UpdatedAt = post.UpdatedAt
	return pr
}

func NewPostsListResponse(posts []models.Post, count int64) *PostListResponse {
	p := new(PostListResponse)
	p.Posts = make([]*PostResponse, 0)
	for _, ls := range posts {
		pr := NewPostResponse(&ls)
		p.Posts = append(p.Posts, pr)
	}
	p.PostCount = count
	return p
}
