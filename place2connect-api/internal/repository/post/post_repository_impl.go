package post

import (
	"errors"

	"github.com/fxfrancky/place2connect-api/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostRepositoryImpl struct {
	DB *gorm.DB
}

// GetPostById implements PostRepository
func (p *PostRepositoryImpl) GetPostById(id int) (*models.Post, error) {
	var post models.Post
	err := p.DB.Preload(clause.Associations).First(&post, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &post, err

}

// CommentPost implements PostRepository
func (p *PostRepositoryImpl) AddCommentToPost(user *models.User, post *models.Post, commentDescription string) error {
	var comment models.Comment
	comment.CommentDescription = commentDescription
	comment.User = *user
	comment.UserID = user.ID
	err := p.DB.Model(post).Association("Comments").Append(&comment)
	if err != nil {
		return err
	}
	err = p.DB.Where(comment.ID).Preload("Post").First(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

// CreatePost implements PostRepository
func (p *PostRepositoryImpl) CreatePost(user *models.User, postDescription string, picturePath string) error {
	var post models.Post
	post.FirstName = user.FirstName
	post.LastName = user.LastName
	post.Location = user.Location
	post.Description = postDescription
	post.PicturePath = picturePath
	post.UserPicturePath = user.PicturePath
	post.User = *user

	result := p.DB.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FeedPosts implements PostRepository
func (p *PostRepositoryImpl) FeedPosts(offset int, limit int) ([]models.Post, int64, error) {
	var (
		posts []models.Post
		count int64
	)

	p.DB.Model(&posts).Count(&count)
	p.DB.Preload(clause.Associations).Offset(offset).Limit(limit).Find(&posts)

	return posts, count, nil
}

// GetUserPosts implements PostRepository
func (p *PostRepositoryImpl) GetUserPosts(userID int, offset int, limit int) ([]models.Post, int64, error) {

	var (
		posts []models.Post
		count int64
	)

	// preload all associations
	// db.Preload(clause.Associations).Find(&users)
	err := p.DB.Where("user_id = ?", userID).Preload(clause.Associations).Offset(offset).Limit(limit).Find(&posts).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	count = int64(len(posts))
	return posts, count, nil
}

// LikePost implements PostRepository
func (p *PostRepositoryImpl) AddLikeToPost(user *models.User, post *models.Post, isLikable bool) error {
	var like models.Like
	like.IsLikablePost = isLikable
	like.User = *user
	like.UserID = user.ID
	err := p.DB.Model(post).Association("Likes").Append(&like)
	if err != nil {
		return err
	}
	err = p.DB.Where(like.ID).Preload("Post").First(&like).Error
	if err != nil {
		return err
	}
	return nil
}

// RemoveLikeToPost implements UserRepository
func (p *PostRepositoryImpl) RemoveLikeToPost(user *models.User, post *models.Post) error {
	var likes = post.Likes

	var err error
	for _, like := range likes {
		if like.UserID == user.ID {
			err = p.DB.Model(&post).Association("Likes").Delete(&like, "CASCADE")
			if err != nil {
				return err
			}
		}

	}

	return err
}

func NewPostRepositoryImpl(DB *gorm.DB) PostRepository {
	return &PostRepositoryImpl{DB: DB}
}
