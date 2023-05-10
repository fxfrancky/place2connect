package post

import "github.com/fxfrancky/place2connect-api/internal/models"

type PostRepository interface {

	// createPost by (userId, postDescription, picturePath)
	CreatePost(user *models.User, postDescription string, picturePath string) error

	// get All Post (feedPosts) getFeedPosts
	FeedPosts(offset, limit int) ([]models.Post, int64, error)

	// get All userID Posts  getUserPosts (userID)
	GetUserPosts(userID int, offset, limit int) ([]models.Post, int64, error)

	// Like a post by (userID, postID, isLikable )
	AddLikeToPost(user *models.User, post *models.Post, isLikable bool) error

	// RemoveLikeToPost implements UserRepository
	RemoveLikeToPost(user *models.User, post *models.Post) error

	// Comment a post by (userID, postID, CommentDescription )
	AddCommentToPost(user *models.User, post *models.Post, commentDescription string) error

	GetPostById(id int) (*models.Post, error)
}
