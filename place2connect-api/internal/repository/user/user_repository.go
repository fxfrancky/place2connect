package user

import "github.com/fxfrancky/place2connect-api/internal/models"

type UserRepository interface {
	// User Repository
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	RegisterUser(user models.User) error
	UpdateUser(user *models.User) error
	AllUsers(offset, limit int) ([]models.User, int64, error)
	DeleteUser(user *models.User) error
	// getUserFriends List By User ID
	GetUserFriends(userID int, offset, limit int) ([]models.User, int64, error)
	// addUserFriend   By User ID And Friend Id
	AddUserFriend(user *models.User, userFriend *models.User) error
	// removeUserFriend   By User ID And Friend Id
	RemoveUserFriend(user *models.User, userFriend *models.User) error
}
