package handlers

import (
	"time"

	"github.com/fxfrancky/place2connect-api/internal/models"
)

// User Response
type UserResponse struct {
	ID            uint            `json:"id,omitempty"`
	LastName      string          `json:"lastName,omitempty"`
	FirstName     string          `json:"firstName,omitempty"`
	Email         string          `json:"email,omitempty"`
	Role          string          `json:"role,omitempty"`
	Photo         string          `json:"photo,omitempty"`
	Provider      string          `json:"provider"`
	IsAdmin       bool            `json:"isAdmin"`
	PicturePath   string          `json:"picturePath"`
	Friends       []*UserResponse `json:"friends"`
	Location      string          `json:"location"`
	Occupation    string          `json:"occupation"`
	Verified      bool            `json:"isVerified"`
	ViewedProfile uint            `json:"viewedProfile"`
	Impressions   uint            `json:"impressions"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserListResponse struct {
	Users      []*UserResponse `json:"users"`
	UsersCount int64           `json:"usersCount"`
}

func NewUserResponse(u *models.User) *UserResponse {
	friends := NewUserListResponse(u.Friends, int64(len(u.Friends)))
	usr := new(UserResponse)
	usr.ID = u.ID
	usr.LastName = u.LastName
	usr.FirstName = u.FirstName
	usr.Email = u.Email
	usr.Photo = u.Photo
	usr.IsAdmin = u.IsAdmin
	usr.Role = u.Role
	usr.Provider = u.Provider
	usr.Verified = u.Verified
	usr.PicturePath = u.PicturePath
	usr.Friends = friends.Users
	usr.Location = u.Location
	usr.Occupation = u.Occupation
	usr.ViewedProfile = u.ViewedProfile
	usr.Impressions = u.Impressions
	usr.CreatedAt = u.CreatedAt
	usr.UpdatedAt = u.UpdatedAt

	return usr
}

func NewUserListResponse(users []models.User, count int64) *UserListResponse {
	u := new(UserListResponse)
	u.Users = make([]*UserResponse, 0)
	for _, us := range users {
		ur := NewUserResponse(&us)
		u.Users = append(u.Users, ur)
	}
	u.UsersCount = count
	return u
}

func FilterUserRecord(user *models.User) UserResponse {
	friends := NewUserListResponse(user.Friends, int64(len(user.Friends)))
	return UserResponse{
		ID:            user.ID,
		LastName:      user.LastName,
		FirstName:     user.FirstName,
		Email:         user.Email,
		Role:          user.Role,
		Photo:         user.Photo,
		Provider:      user.Provider,
		IsAdmin:       user.IsAdmin,
		PicturePath:   user.PicturePath,
		Friends:       friends.Users,
		Location:      user.Location,
		Occupation:    user.Occupation,
		ViewedProfile: user.ViewedProfile,
		Verified:      user.Verified,
		Impressions:   user.Impressions,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
