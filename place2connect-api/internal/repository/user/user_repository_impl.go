package user

import (
	"errors"
	"strings"

	"github.com/fxfrancky/place2connect-api/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// AllUsers implements UserRepository
func (u *UserRepositoryImpl) AllUsers(offset int, limit int) ([]models.User, int64, error) {
	var (
		users []models.User
		count int64
	)
	u.DB.Model(&users).Count(&count)
	u.DB.Offset(offset).Limit(limit).Find(&users)
	return users, count, nil
}

// DeleteUser implements UserRepository
func (u *UserRepositoryImpl) DeleteUser(user *models.User) error {
	// Delete a user with his associations
	return u.DB.Select(clause.Associations).Delete(user).Error
}

// GetUserByEmail implements UserRepository
func (u *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.DB.Preload("Friends").First(&user, "email = ?", strings.ToLower(email)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

	}

	return &user, err
}

// GetUserById implements UserRepository
func (u *UserRepositoryImpl) GetUserById(id int) (*models.User, error) {
	var user models.User
	err := u.DB.Preload("Friends").First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, err

}

// RegisterUser implements UserRepository
func (u *UserRepositoryImpl) RegisterUser(user models.User) error {
	result := u.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateUser implements UserRepository
func (u *UserRepositoryImpl) UpdateUser(user *models.User) error {
	result := u.DB.Model(&user).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// AddUserFriend implements UserRepository
func (u *UserRepositoryImpl) AddUserFriend(user *models.User, userFriend *models.User) error {
	err := u.DB.Model(user).Association("Friends").Append(userFriend)
	if err != nil {
		return err
	}
	return u.DB.Where(userFriend.ID).First(userFriend).Error
}

// GetUserFriends implements UserRepository
func (u *UserRepositoryImpl) GetUserFriends(userID int, offset int, limit int) ([]models.User, int64, error) {

	var (
		user  models.User
		count int64
	)

	// preload all associations
	// db.Preload(clause.Associations).Find(&users)
	err := u.DB.Where("id = ?", userID).Preload("Friends").Offset(offset).Limit(limit).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	count = int64(len(user.Friends))
	return user.Friends, count, nil
}

// RemoveUserFriend implements UserRepository
func (u *UserRepositoryImpl) RemoveUserFriend(user *models.User, userFriend *models.User) error {
	err := u.DB.Model(&user).Association("Friends").Delete(userFriend)

	return err
}

func NewUserRepositoryImpl(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}
